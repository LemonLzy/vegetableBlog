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

	engine.Use(getCorsConf())
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
		api.POST("/user/modify_pw", handler.UserModifyPwHandler)
		api.POST("/user/modify", handler.UserModifyHandler)
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

func getCorsConf() gin.HandlerFunc {
	// 注意JWT认证时，放行请求头Authorization，避免跨域问题
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Language", "Content-Type"},
		AllowCredentials: false,
	})
}
