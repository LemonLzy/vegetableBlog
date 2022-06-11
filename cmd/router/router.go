package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/cmd/router/handler"
	"net/http"
	"os"
)

func Register(engine *gin.Engine) {
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Println(err)
		return
	}

	indexPath, staticPath := getStaticPath()
	engine.LoadHTMLFiles(indexPath)
	engine.StaticFS("./static", http.Dir(staticPath))
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "baidu.com",
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

func getStaticPath() (string, string) {
	projectPath, _ := os.Getwd()
	indexPath := projectPath + "/web/index.html"
	staticPath := projectPath + "/web/static"
	fmt.Println(projectPath, indexPath, staticPath)
	return indexPath, staticPath
}
