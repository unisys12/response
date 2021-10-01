package auth

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"

	"github.com/Masterminds/sprig"
	"github.com/davecgh/go-spew/spew"

	"github.com/volatiletech/authboss/v3"
)

const (
	renderFileNotFound = `<p>Sorry, the auth page could not be compiled. Either the <b>base.html.tmpl</b> or <b>%s.html.tmpl</b> file is missing.</p>`
	renderCompileError = `<p>Sorry, the auth page could not be compiled.<br /><br /><pre>%s</pre></p>`
	renderContentType  = "text/html"
)

var (
	errPatternMatchesNoFiles = errors.New("pattern matches no files")
)

//go:embed templates
var authPagesEmbeddedFS embed.FS

func (s *Service) newAuthPagesRenderer() (authboss.Renderer, error) {
	if s.config.UIDirAuthPages == "" {
		return &authPagesRenderer{
			pageFS:    authPagesEmbeddedFS,
			templates: map[string]*template.Template{},
		}, nil
	}

	switch s.config.DeveloperDisableCachingAuthPages {
	case true:
		_, _ = fmt.Fprintf(os.Stdout, "> Rendering Auth Pages from %q. %s\n", s.config.UIDirAuthPages, color.RedString("Template caching is disabled."))
	default:
		_, _ = fmt.Fprintf(os.Stdout, "> Rendering Auth Pages from %q. Templates are cached.\n", s.config.UIDirAuthPages)
	}

	return &authPagesRenderer{
		pageFS:       os.DirFS(s.config.UIDirAuthPages),
		templates:    map[string]*template.Template{},
		providers:    s.providers,
		disableCache: s.config.DeveloperDisableCachingAuthPages,
	}, nil
}

type authPagesRenderer struct {
	pageFS       fs.FS
	disableCache bool
	providers    []*provider
	templates    map[string]*template.Template
	templateLock sync.RWMutex
}

func (a *authPagesRenderer) funcs() template.FuncMap {
	funcs := sprig.FuncMap()
	funcs["htmlAttr"] = func(attrs ...string) template.HTMLAttr {
		var str string
		for _, attr := range attrs {
			parts := strings.Split(attr, "=")
			if len(parts) == 1 {
				str = str + fmt.Sprintf("%s ", parts[0])

				continue
			}

			str = str + fmt.Sprintf("%s=%q ", parts[0], parts[1])
		}

		return template.HTMLAttr(str)
	}

	return funcs
}

func (a *authPagesRenderer) loadSinglePage(name string) (*template.Template, error) {
	tmpl, err := template.New("page").
		Funcs(a.funcs()).
		ParseFS(a.pageFS, "partials/*.html.tmpl", "_base.html.tmpl", fmt.Sprintf("%s.html.tmpl", name))

	if err != nil {
		return nil, fmt.Errorf("template.New...ParseFS: %w", err)
	}

	if tmpl == nil {
		return nil, fmt.Errorf("nil template %q", name)
	}

	return tmpl, nil
}

func (a *authPagesRenderer) Load(names ...string) error {
	a.templateLock.Lock()
	defer a.templateLock.Unlock()

	for _, name := range names {
		tmpl, err := a.loadSinglePage(name)
		if err != nil {
			return fmt.Errorf("load auth page template: %w", err)
		}

		if !a.disableCache {
			a.templates[name] = tmpl
		}
	}

	return nil
}

func (a *authPagesRenderer) Render(ctx context.Context, page string, data authboss.HTMLData) ([]byte, string, error) {
	a.templateLock.RLock()
	defer a.templateLock.RUnlock()

	var err error
	var buf bytes.Buffer
	var tmpl = a.templates[page]

	if a.disableCache {
		tmpl, err = a.loadSinglePage(page)
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "pattern matches no files"):
				buf.WriteString(fmt.Sprintf(renderFileNotFound, page))
				return buf.Bytes(), renderContentType, err
			default:
				buf.WriteString(fmt.Sprintf(renderCompileError, err.Error()))
				return buf.Bytes(), renderContentType, err
			}
		}
	}

	data = data.Merge(authboss.HTMLData{
		"providers": a.providers,
		"csrfToken": "",
	})

	spew.Dump(data)

	if tmpl == nil {
		buf.WriteString(fmt.Sprintf(renderFileNotFound, page))
		return buf.Bytes(), renderContentType, err
	}

	if err := tmpl.Funcs(a.funcs()).ExecuteTemplate(&buf, "base", data); err != nil {
		buf.WriteString(fmt.Sprintf(renderCompileError, err.Error()))
		return buf.Bytes(), renderContentType, err
	}

	return buf.Bytes(), renderContentType, err
}
