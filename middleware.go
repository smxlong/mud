package mud

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// auth is a middleware that authenticates the request by looking for a JWT
// in the Authorization header, or in a cookie called "token".
func (s *Service) auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	// if that's empty, look at the cookie
	if authHeader == "" {
		authHeader, _ = c.Cookie("token")
		if authHeader != "" {
			s.l.Debugw("found token in cookie")
			authHeader = "Bearer " + authHeader
		}
	}
	if authHeader == "" {
		c.Next()
		return
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	player, err := verifyPlayer(c, s.entcli, s.JWTAudience, s.JWTIssuer, []byte(s.JWTSigningKey), token)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	c.Set("player", player)
	c.Next()
}

// logRequest logs the request
func (s *Service) logRequest(c *gin.Context) {
	l := s.l.With("middleware", "logRequest")
	start := time.Now()
	l.Infow("request started", "method", c.Request.Method, "path", c.Request.URL.Path)
	defer func() {
		elapsed := time.Since(start)
		l.Infow("request completed", "method", c.Request.Method, "path", c.Request.URL.Path, "status", c.Writer.Status(), "elapsed", elapsed)
	}()
	c.Next()
}

// requirePlayer is a middleware that requires a player to be logged in
func (s *Service) requirePlayer(c *gin.Context) {
	l := s.l.With("middleware", "requirePlayer")
	l.Debugw("checking for player")
	player, ok := c.Get("player")
	if !ok {
		l.Errorw("no player found")
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	l.Debugw("player found", "player", player)
	c.Set("player", player)
	c.Next()
}

// setupMiddleware sets up the middleware for the service
func (s *Service) setupMiddleware() {
	s.g.Use(
		s.logRequest,
		gin.Recovery(),
		s.auth,
	)
}
