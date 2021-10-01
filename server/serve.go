package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func (s *server) Addr() string {
	return s.config.HTTPBindAddress
}

func (s *server) HTTPProtocol() string {
	if s.HTTPTLSEnabled() {
		return "https"
	}

	return "http"
}

func (s *server) NewActor(ctx context.Context) (func() error, func(error)) {
	ctx, cancel := context.WithCancel(ctx)

	runFunc := func() error {
		return s.Serve(ctx)
	}

	intFunc := func(err error) {
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			s.logger.Error().
				Err(err).
				Msg("failed to Shutdown server")
		}
	}

	return runFunc, intFunc
}

func (s *server) Serve(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	case err := <-s.serveHTTP():
		return err
	}
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func (s *server) serveHTTP() <-chan error {
	if !s.config.HideServing {
		switch s.HTTPProtocol() {
		case "http":
			_, _ = fmt.Fprintf(os.Stdout, "> Serving %s TLS on %s://%s. You should use TLS in production!\n", color.RedString("without"), s.HTTPProtocol(), net.JoinHostPort(s.Addr(), strconv.Itoa(s.config.HTTPPort)))
		default:
			_, _ = fmt.Fprintf(os.Stdout, "> Serving %s TLS on %s://%s\n", color.GreenString("with"), s.HTTPProtocol(), net.JoinHostPort(s.Addr(), strconv.Itoa(s.config.HTTPTLSPort)))
		}

		if s.isDevProxyEnabled() {
			_, _ = fmt.Fprintf(os.Stdout, color.YellowString("> Hey Developer! Response UI is being proxied from http://127.0.0.1:%s\n"), strconv.Itoa(s.config.DeveloperUIProxyPort))
		}

		if !s.isServingEmbedded() {
			_, _ = fmt.Fprintf(os.Stdout, color.YellowString("> Response UI is being served from %s\n"), s.config.UIDirConsoleDist)
		}
	}

	if s.config.DeveloperProfiling {
		_, _ = fmt.Fprintf(os.Stdout, "\n%s\n%s\n", color.RedString("!!!!! Profiling endpoints are enabled. !!!!!"), color.BlueString("Profiling is a developer tool that should only be enabled when requested or if you know what you're doing. Be careful enabling this in production."))
	}

	err := make(chan error, 1)
	switch s.HTTPTLSMode() {
	case TLSModeStatic:
		go s.serveHTTPWithStaticTLS(err)
	case TLSModeAuto:
		go s.serveHTTPWithAutoTLS(err)
	}

	go s.serveHTTPWithoutTLS(err)

	return err
}

func (s *server) serveHTTPWithoutTLS(err chan<- error) {
	s.http = &http.Server{
		Addr:    net.JoinHostPort(s.Addr(), strconv.Itoa(s.config.HTTPPort)),
		Handler: s.routes,
	}

	err <- s.http.ListenAndServe()
}
