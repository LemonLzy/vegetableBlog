package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/oss"
	"net/http"
)

func OssPostHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		loc := oss.Put(file)
		c.JSON(http.StatusOK, gin.H{
			"location": loc,
		})
	}
}
