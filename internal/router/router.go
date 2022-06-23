package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handler2 "github.com/lemonlzy/vegetableBlog/internal/router/handler"
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
		api.POST("/sign_up", handler2.UserSignUpHandler)
		api.POST("/sign_in", handler2.UserSignInHandler)
	}

	{
		api.GET("/article/:id", handler2.ArticleDetailHandler)
		api.GET("/articles", handler2.ArticleListHandler)
		api.POST("/article", handler2.ArticleCreateHandler)
		api.PUT("/article/:id", handler2.ArticleUpdateHandler)
	}

	{
		api.POST("/user", handler2.UserCreateHandler)
		api.PUT("/user/:id", handler2.UserUpdateHandler)
	}

	oss := engine.Group("/oss")
	{
		oss.POST("/upload", handler2.OssPostHandler)
	}
}

func getStaticPath() (string, string) {
	projectPath, _ := os.Getwd()
	indexPath := projectPath + "/web/index.html"
	staticPath := projectPath + "/web/static"
	fmt.Println(projectPath, indexPath, staticPath)
	return indexPath, staticPath
}
