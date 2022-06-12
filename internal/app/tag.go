package app

type Tag struct {
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	Name      string `json:"name,omitempty" gorm:"type:varchar(30);comment:标签名称"`
	CreatedAT int64  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAT int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
