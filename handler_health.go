package mud

import "github.com/gin-gonic/gin"

// healthz is the health check endpoint
func (s *Service) healthz(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
