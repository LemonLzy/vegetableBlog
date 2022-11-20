package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/oss"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/resp"
)

func OssPostHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		loc := oss.Put(file)
		resp.ResponseSuccess(c, gin.H{
			"location": loc,
		})
	}
}
