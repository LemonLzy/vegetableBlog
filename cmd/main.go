package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/cmd/vBlog"
)

func main() {
	err := vBlog.Init()
	if err != nil {
		fmt.Println(err)
	}

	engine := gin.New()
	vBlog.Register(engine)
	err = engine.Run(":" + vBlog.Conf.Port)
	if err != nil {
		fmt.Println(err)
	}
}
