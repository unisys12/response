package server

import (
	"net/http"
	"os"
	"path"
)

func (s *server) isServingEmbedded() bool {
	if s.config.UIDirConsoleDist == "" {
		return true
	}

	_, err := os.Stat(path.Join(s.config.UIDirConsoleDist, "index.html"))
	if os.IsNotExist(err) {
		s.logger.Error().
			Err(err).
			Msg("ui_dir directory does not contain index.html, the ui_dir is being ignored")

		return true
	}

	if os.IsPermission(err) {
		s.logger.Error().
			Err(err).
			Msg("the user running Response does not have permission for the ui_dir directory, the ui_dir is being ignored")

		return true
	}

	if err != nil {
		s.logger.Error().
			Err(err).
			Msg("there was a problem with the path specified in ui_dir, the ui_dir is being ignored")

		return true
	}

	return false
}

func (s *server) handleConsoleRoutes() http.HandlerFunc {
	//var servingFS fs.FS
	//switch {
	//case s.isServingEmbedded():
	//	servingFS = ui.UIAssetsFS
	//default:
	//	servingFS = os.DirFS(s.config.UIDirConsoleDist)
	//}
	//
	//return func(wr http.ResponseWriter, r *http.Request) {
	//	p := path.Clean(path.Join("dist", c.Path()))
	//
	//	_, err := servingFS.Open(p)
	//	if !os.IsNotExist(err) {
	//		http.FileServer(http.FS(servingFS)).
	//			ServeHTTP(c.Response().Writer, c.Request())
	//
	//		return nil
	//	}
	//
	//	index, err := servingFS.Open(path.Join("dist", "index.html"))
	//	if os.IsNotExist(err) {
	//		return c.String(http.StatusInternalServerError, "Sorry, a problem occurred")
	//	}
	//
	//	indexBytes, err := ioutil.ReadAll(index)
	//	if err != nil {
	//		return c.String(http.StatusInternalServerError, "Sorry, a problem occurred")
	//	}
	//
	//	return c.Blob(http.StatusOK, http.DetectContentType(indexBytes), indexBytes)
	//}

	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
