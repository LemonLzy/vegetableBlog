package app

import (
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
	"gorm.io/gorm"
)

type Article struct {
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	TagID     int    `json:"tag_id,omitempty" gorm:"comment:标签ID;not null"`
	Status    int8   `json:"status,omitempty" gorm:"comment:文章状态：0-草稿 1-已发布 2-已删除;default:0"`
	ArticleID int64  `json:"article_id,omitempty" gorm:"index:idx_article_id,unique;comment:文章ID，便于url访问"`
	Title     string `json:"title,omitempty" gorm:"type:varchar(30);comment:文章标题;not null"`
	Path      string `json:"path,omitempty" gorm:"type:varchar(30);comment:文章路径;not null"`
	Summary   string `json:"summary,omitempty" gorm:"type:varchar(100);comment:文章摘要;not null"`
	Content   string `json:"content,omitempty" gorm:"comment:文章内容;not null"`
	RichText  string `json:"rich_text,omitempty" gorm:"comment:文章富文本内容;not null"`
	Cover     string `json:"cover,omitempty" gorm:"type:varchar(100);comment:封面图片"`
	CreatedAT int64  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAT int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

// CreateArticle 创建文章
func CreateArticle(a *Article) error {
	// 生成文章ID
	a.ArticleID = snowflake.GenID()
	// 数据库创建
	if err := DB.Debug().Create(a).Error; err != nil {
		// 日志记录
		return err
	}
	return nil
}

// GetArticleDetail 根据文章id获取文章详情
func GetArticleDetail(articleID int64) (*Article, error) {
	a := new(Article)
	if err := DB.Debug().Where("article_id = ?", articleID).First(&a).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return a, nil
}

// GetArticleList 获取所有文章
func GetArticleList(page, size int64) ([]*Article, error) {
	var articles []*Article
	if err := DB.Debug().Where("status != ?", 2).Offset(int((page - 1) * size)).Order("updated_at desc").Limit(int(size)).Find(&articles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return articles, nil
}

// UpdateArticleByID 更新文章信息
func UpdateArticleByID(articleID int64, a *Article) error {
	if err := DB.Debug().Where("article_id = ?", articleID).Updates(a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return nil
}
