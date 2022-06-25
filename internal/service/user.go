package service

import (
	"github.com/lemonlzy/vegetableBlog/api"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	errCode "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/utils"
)

func SignUp(psu *api.ParamSignUp) error {
	// 判断用户名是否重复
	if ok, err := app.GetUserByName(psu.Username); ok || err != nil {
		return err
	}

	if psu.Password != psu.RePassword {
		return errCode.NewClientError(errCode.CodeUserPwNotEqual, "两次输入的密码不相等")
	}
	// 生成用户唯一标识ID
	userID := snowflake.GenID()
	// 密码加密
	hashPw := utils.BcryptPw(psu.Password)
	user := &app.User{
		Username: psu.Username,
		Password: hashPw,
		Nickname: utils.GenRandNickName(),
		UserID:   userID,
	}
	// 存储
	if err := app.CreateUser(user); err != nil {
		return err
	}
	return nil
}
