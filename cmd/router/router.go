package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/cmd/router/handler"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "成功",
		})
	})

	v1 := engine.Group("/api/v1")
	{
		v1.GET("/article/:id", handler.ArticleDetailHandler)
		v1.GET("/articles", handler.ArticleListHandler)
		v1.POST("/article", handler.CreateArticleHandler)
		v1.PUT("/article/:id", handler.UpdateArticleHandler)
	}

	oss := engine.Group("/oss")
	{
		oss.POST("/upload", handler.OssPostHandler)
	}
}
