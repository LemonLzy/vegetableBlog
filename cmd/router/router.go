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

	api := engine.Group("/api/")
	{
		api.POST("/sign_up", handler.UserSignUpHandler)
		api.POST("/sign_in", handler.UserSignInHandler)
	}

	{
		api.GET("/article/:id", handler.ArticleDetailHandler)
		api.GET("/articles", handler.ArticleListHandler)
		api.POST("/article", handler.ArticleCreateHandler)
		api.PUT("/article/:id", handler.ArticleUpdateHandler)
	}

	{
		api.POST("/user", handler.UserCreateHandler)
		api.PUT("/user/:id", handler.UserUpdateHandler)
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
