package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type TraceLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (t TraceLogWriter) Write(p []byte) (int, error) {
	if n, err := t.body.Write(p); err != nil {
		return n, err
	}
	return t.ResponseWriter.Write(p)
}

func TraceLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &TraceLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		log.Printf("access log: method: %s, code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
