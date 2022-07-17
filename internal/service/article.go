package service

import (
	"github.com/lemonlzy/vegetableBlog/internal/app"
	"github.com/lemonlzy/vegetableBlog/pkg"
)

// ArticleCount 根据userID统计对应的文章信息
func ArticleCount(userID int64) ([]*app.Article, error) {
	// 判断用户是否存在
	if ok, err := app.CheckUserByID(userID); !ok || err != nil {
		return nil, err
	}

	// 获取分页信息
	page, size := pkg.GetPageInfo()

	// 查询文章信息
	articles, err := app.GetArticlesByUserID(userID, page, size)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
