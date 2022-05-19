package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/cmd/vegetableBlog"
)

func main() {
	engine := gin.New()
	vegetableBlog.Register(engine)
	engine.Run()
}
