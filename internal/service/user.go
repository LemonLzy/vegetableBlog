package service

import (
	"github.com/lemonlzy/vegetableBlog/api"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
)

func SignUp(psu *api.ParamSignUp) error {
	// 判断用户名是否重复
	if ok, err := app.GetUserByName(psu.Username); ok || err != nil {
		return err
	}
	// 生成用户唯一标识ID
	userID := snowflake.GenID()
	// 密码加密

	user := &app.User{
		Username: psu.Username,
		Password: psu.Password,
		Nickname: "",
		UserID:   userID,
	}
	// 存储
	if err := app.CreateUser(user); err != nil {
		return err
	}
	return nil
}
