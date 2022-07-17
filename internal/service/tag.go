package service

import (
	"github.com/lemonlzy/vegetableBlog/internal/app"
)

// TagCount 获取所有的tag名称
func TagCount() ([]string, error) {
	// 获取tag_name的列表
	tagNames, err := app.GetTags()
	if err != nil {
		return nil, err
	}

	return tagNames, nil
}
