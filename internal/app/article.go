package app

import (
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
	"github.com/lemonlzy/vegetableBlog/internal/server/mysql"
)

type Article struct {
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	TagID     int    `json:"tag_id,omitempty" gorm:"comment: 标签ID"`
	Status    int8   `json:"status,omitempty" gorm:"comment: 文章状态：0-草稿 1-已发布 2-已删除"`
	ArticleID int64  `json:"article_id,omitempty" gorm:"index:idx_article_id,unique;comment: 文章ID，便于url访问"`
	Title     string `json:"title,omitempty" gorm:"comment: 文章标题"`
	Path      string `json:"path,omitempty" gorm:"comment: 文章路径"`
	Summary   string `json:"summary,omitempty" gorm:"comment: 文章摘要"`
	Content   string `json:"content,omitempty" gorm:"comment: 文章内容"`
	RichText  string `json:"rich_text,omitempty" gorm:"comment: 文章富文本内容"`
	CreatedAT int64  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAT int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func CreateArticle(a *Article) error {
	// 生成文章ID
	a.ArticleID = snowflake.GenID()
	// 数据库创建
	if err := mysql.DB.Debug().Create(a).Error; err != nil {
		// 日志记录
		return err
	}
	return nil
}
