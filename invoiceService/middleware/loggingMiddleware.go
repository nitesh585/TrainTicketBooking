package middleware

import (
	"invoiceService/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		log.WithFields(logrus.Fields{
			"method":          c.Request.Method,
			"path":            c.Request.URL.Path,
			"clientIP":        c.ClientIP(),
			"clientUserAgent": c.Request.UserAgent(),
			"dataLength":      c.Writer.Size(),
			"referer":         c.Request.Referer(),
			"status":          c.Writer.Status(),
			"latency_us":      time.Since(start).Microseconds(),
		}).Info()
	}
}
