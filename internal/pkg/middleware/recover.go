package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				log.Printf(s, err)
				c.Abort()
			}
		}()
		c.Next()
	}
}
