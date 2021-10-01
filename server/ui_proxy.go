package server

func (s *server) isDevProxyEnabled() bool {
	return s.config.DeveloperUIProxyPort != 0
}

func (s *server) registerDevProxyRoute() {
	if s.config.DeveloperUIProxyPort != 0 {
		//var consoleTargets = []*middleware.ProxyTarget{
		//	{
		//		URL: &url.URL{
		//			Scheme: "http",
		//			Host:   net.JoinHostPort("127.0.0.1", strconv.Itoa(s.config.DeveloperUIProxyPort)),
		//		},
		//	},
		//}

		//g := s.routes.Group("/*")
		//g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(consoleTargets)))
	}
}
