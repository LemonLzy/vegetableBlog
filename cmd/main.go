package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/cmd/vegetableBlog"
)

func main() {
	err := vegetableBlog.Init()
	if err != nil {
		fmt.Println(err)
	}

	engine := gin.New()
	vegetableBlog.Register(engine)
	err = engine.Run(":" + vegetableBlog.Conf.Port)
	if err != nil {
		fmt.Println(err)
	}
}
