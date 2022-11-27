package service

import (
	"github.com/lemonlzy/vegetableBlog/internal/app"
	errCode "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/utils"
)

// ArticleCount 根据userID统计对应的文章信息
func ArticleCount(userID, page, size int64) ([]*app.Article, error) {
	// 判断用户是否存在
	if ok, err := app.CheckUserByID(userID); !ok || err != nil {
		return nil, err
	}

	// 查询文章信息
	articles, err := app.GetArticlesByUserID(userID, page, size)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// CreateArticle 创建文章
func CreateArticle(a *app.Article) error {
	article := &app.Article{
		TagID:     a.TagID,
		Status:    a.Status,
		UserID:    a.UserID,
		ArticleID: snowflake.GenIDStr(),
		Title:     a.Title,
		Path:      a.Path,
		Summary:   utils.SubStr(a.Content, 0, 100), // 摘要：截取前100个字符
		Content:   a.Content,
		RichText:  a.RichText,
		Cover:     a.Cover,
	}

	err := app.CreateArticle(article)
	if err != nil {
		return err
	}
	return nil
}

// PubArticle 发布文章
func PubArticle(articleID string) error {
	// 判断文章是否存在
	a, err := app.GetArticleByID(articleID)
	if err != nil {
		return err
	}
	// 判断文章状态是否为草稿
	if a.Status != 0 {
		return errCode.NewServerError(errCode.ArticleStatus)
	}
	// 发布文章
	a.Status = 1
	err = app.UpdateArticleByID(articleID, a)
	if err != nil {
		return err
	}
	return nil
}

// UpdateArticle 更新文章
func UpdateArticle(a *app.Article) error {
	return nil
}
