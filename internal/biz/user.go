package biz

type User struct {
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	Username  string `json:"username,omitempty" gorm:"comment: 用户登录名"`
	Password  string `json:"password,omitempty" gorm:"comment: 登录密码"`
	Nickname  string `json:"nickname,omitempty" gorm:"comment: 用户昵称"`
	IsAdmin   bool   `json:"is_admin,omitempty" gorm:"comment: 是否是管理员：0-否 1-是"`
	CreatedAT int64  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAT int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
