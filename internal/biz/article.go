package biz

type Article struct {
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	TagID     int    `json:"tag_id,omitempty"`
	Status    int8   `json:"status,omitempty"`
	ArticleID int64  `json:"article_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Path      string `json:"path,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Content   string `json:"content,omitempty"`
	RichText  string `json:"rich_text,omitempty"`
	CreatedAT int64  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAT int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
