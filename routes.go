package mud

// setupRoutes sets up the routes for the service
func (s *Service) setupRoutes() {
	s.g.GET("/healthz", s.healthz)
	s.g.GET("/auth/healthz", s.requirePlayer, s.healthz)
	s.g.POST("/login", s.login)
	s.g.POST("/logout", s.logout)
}
