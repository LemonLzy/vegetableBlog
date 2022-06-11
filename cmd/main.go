package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/cmd/router"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
	"github.com/lemonlzy/vegetableBlog/internal/server/conf"
	"github.com/lemonlzy/vegetableBlog/internal/server/mysql"
	"net/http"
	"os"
)

func main() {
	err := conf.Init()
	if err != nil {
		fmt.Println(err)
	}

	err = mysql.Init(conf.Conf.DBConfig)
	if err != nil {
		fmt.Println(err)
	}

	err = snowflake.Init(conf.Conf.SinceTime, 1)
	if err != nil {
		fmt.Printf("init snowflake failed. err:%v\n", err)
		return
	}

	engine := gin.New()
	router.Register(engine)
	fmt.Println(os.Getwd())
	engine.LoadHTMLFiles("./web/index.html")
	engine.StaticFS("./static", http.Dir("./web/static/"))
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "baidu.com",
		})
	})

	err = engine.Run(":" + conf.Conf.Port)
	if err != nil {
		fmt.Println(err)
	}
}
