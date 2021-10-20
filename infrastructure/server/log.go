package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/digilibP/infrastructure/log"
	"github.com/sirupsen/logrus"
)

// Log to file
func LoggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		loggerContext := log.ContextWithLogGroupID(c.Request.Context())

		startTime := time.Now()
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		clientIP := c.ClientIP()
		log.InfoRequest(loggerContext, (logrus.Fields{
			"time":       startTime,
			"client_ip":  clientIP,
			"req_method": reqMethod,
			"req_uri":    reqUri,
		}))

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		statusCode := c.Writer.Status()

		log.InfoResponse(loggerContext, (logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}))
	}
}

// Log to MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Log to ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Logging to MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
