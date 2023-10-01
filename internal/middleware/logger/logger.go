package logger

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("log", log)
		c.Next()
	}
}
