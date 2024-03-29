package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/router/handler"
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

	engine.Use(getCorsFunc())
	api := engine.Group("/api/")
	{
		api.POST("/sign_up", handler.UserSignUpHandler)
		api.POST("/sign_in", handler.UserSignInHandler)
	}

	{
		api.GET("/toc/:userID", handler.UserFontPageHandler)
	}

	{
		api.GET("/article/:articleID", handler.ArticleDetailHandler)
		api.GET("/articles", handler.ArticleListHandler)
		api.GET("/articles/archive", handler.ArticleArchiveHandler)
		api.POST("/article", handler.ArticleCreateHandler)
		api.POST("/article/:id", handler.ArticleUpdateHandler)
		api.POST("/article/pub", handler.ArticlePubHandler)
	}

	{
		api.POST("/user/modify_pw", handler.UserModifyPwHandler)
		api.POST("/user/modify", handler.UserModifyHandler)
		api.POST("/user/del", handler.UserDelHandler)
	}

	{
		// TODO 用户配置相关信息
		//api.POST("/config", handler.UserConfigHandler)
	}

	{
		api.POST("/upload_image", handler.OssPostHandler)
	}
}

func getStaticPath() (string, string) {
	projectPath, _ := os.Getwd()
	indexPath := projectPath + "/web/index.html"
	staticPath := projectPath + "/web/static"
	fmt.Println(projectPath, indexPath, staticPath)
	return indexPath, staticPath
}

func getCorsFunc() gin.HandlerFunc {
	// 注意JWT认证时，放行请求头Authorization，避免跨域问题(这里放行了所有请求头)
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Language", "Content-Type"},
		AllowCredentials: false,
	})
}
