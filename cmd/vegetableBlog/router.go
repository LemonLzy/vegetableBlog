package vegetableBlog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "成功",
		})
	})
}
