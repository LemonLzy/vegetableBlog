package pkg

import (
	"strconv"
)

// 固定每页10条记录，不可更改
const pageSize = "10"

func GetPageInfo() (int64, int64) {
	// 获取分页参数
	pageStr := "1"
	sizeStr := pageSize

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
