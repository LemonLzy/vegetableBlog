package pkg

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageInfo(c *gin.Context) (int64, int64) {
	// 获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
