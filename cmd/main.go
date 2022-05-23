package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/server/conf"
	"github.com/lemonlzy/vegetableBlog/internal/server/http"
	"github.com/lemonlzy/vegetableBlog/internal/server/mysql"
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

	engine := gin.New()
	http.Register(engine)
	err = engine.Run(":" + conf.Conf.Port)
	if err != nil {
		fmt.Println(err)
	}
}
