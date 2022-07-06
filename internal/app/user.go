package app

import (
	errCode "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"gorm.io/gorm"
)

type User struct {
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	Username  string `json:"username,omitempty" gorm:"type:varchar(30);uniqueIndex;comment:用户登录名" binding:"required"`
	Password  string `json:"password,omitempty" gorm:"type:varchar(64);comment:登录密码" binding:"required"`
	Nickname  string `json:"nickname,omitempty" gorm:"type:varchar(30);comment:用户昵称"`
	IsAdmin   bool   `json:"is_admin,omitempty" gorm:"default:0;comment:是否是管理员：0-否 1-是;"`
	UserID    int64  `json:"user_id,omitempty" gorm:"comment:用户唯一ID"`
	CreatedAT int64  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAT int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAT int64  `json:"deleted_at,omitempty" gorm:"autoDeleteTime"`
}

// CreateUser 创建用户
func CreateUser(u *User) error {
	if err := DB.Debug().Create(u).Error; err != nil {
		// 日志记录
		return err
	}
	return nil
}

// UpdateUserByID 更新用户信息
func UpdateUserByID(userID int64, u *User) error {
	err := DB.Debug().Where("user_id = ?", userID).Updates(u).Error
	return err
}

// GetUserByName 根据用户名查找用户
func GetUserByName(name string) (bool, error) {
	var u User
	err := DB.Debug().Select("id").Where("username = ? AND deleted_at = ?", name, 0).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if u.ID > 0 {
		return true, errCode.NewClientError(errCode.CodeUserExist)
	}

	return false, nil
}

// GetUserByID 根据用户ID查找用户
func GetUserByID(userID int64) (bool, error) {
	var u User
	err := DB.Debug().Select("id").Where("user_id = ?  AND deleted_at = ?", userID, 0).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if u.ID > 0 {
		return true, nil
	}

	return false, errCode.NewClientError(errCode.CodeUserNotExist)
}

// GetUserPwByName 根据用户名查找用户密码
func GetUserPwByName(name string) (string, error) {
	var u User
	err := DB.Debug().Select("password").Where("username = ? AND deleted_at = ?", name, 0).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if u.Password != "" {
		return u.Password, nil
	}

	return "", errCode.NewClientError(errCode.CodeUserExist)
}

// GetUserPwByID 根据用户ID查找用户密码
func GetUserPwByID(userID int64) (string, error) {
	var u User
	err := DB.Debug().Select("password").Where("user_id = ? AND deleted_at = ?", userID, 0).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if u.Password != "" {
		return u.Password, nil
	}

	return "", errCode.NewClientError(errCode.CodeUserExist)
}
