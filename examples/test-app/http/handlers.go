package http

import (
	"test-app/events"
	"test-app/metrics"

	"github.com/gin-gonic/gin"
)

func homeHandler(version string) func(c *gin.Context) {
	return func(c *gin.Context) {
		broker, err := events.NewClient("test-app-internal-api-server", "http://default-broker")
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err,
			})
		}
		err = broker.SendEvent(c)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err,
			})
		}
		c.JSON(200, gin.H{
			"app_version": version,
			"git_commit":  "s",
		})
	}
}

func metricsHandler(pe metrics.MetricsService) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler := pe.GetHandler()
		handler(c.Writer, c.Request)
	}
}
