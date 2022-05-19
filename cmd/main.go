package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/server"
)

func main() {
	err := server.Init()
	if err != nil {
		fmt.Println(err)
	}

	engine := gin.New()
	server.Register(engine)
	err = engine.Run(":" + server.Conf.Port)
	if err != nil {
		fmt.Println(err)
	}
}
