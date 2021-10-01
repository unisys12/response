package server

import (
	"net/http/pprof"
)

func (s *server) registerDeveloperRoutes() {

	// pprof
	if s.config.DeveloperProfiling {
		s.routes.Get("/debug/pprof/*", pprof.Index)
	}
}
