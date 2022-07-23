package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 固定每页10条记录，不可更改
const pageSize = 10

// GetPageInfo 获取分页参数
func GetPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	size := pageSize

	page, err := strconv.ParseInt(pageStr, 10, 10)
	if err != nil {
		page = 1
	}

	return page, int64(size)
}

// GetUserID 从上下文cookies获取userID
func GetUserID(c *gin.Context) int64 {
	cookieID, err := c.Cookie("user_id")
	if err != nil {
		log.Fatal(err)
	}

	userID, err := strconv.ParseInt(cookieID, 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	return userID
}
