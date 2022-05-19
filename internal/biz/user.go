package biz

type User struct {
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Role      string `json:"role,omitempty"`
	CreatedAT int64  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAT int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
