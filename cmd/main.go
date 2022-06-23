package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
	"github.com/lemonlzy/vegetableBlog/internal/router"
	"github.com/lemonlzy/vegetableBlog/internal/server/conf"
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

	err = snowflake.Init(conf.Conf.SinceTime, 1)
	if err != nil {
		fmt.Printf("init snowflake failed. err:%v\n", err)
		return
	}

	engine := gin.New()
	router.Register(engine)

	err = engine.Run(":" + conf.Conf.Port)
	if err != nil {
		fmt.Println(err)
	}
}
