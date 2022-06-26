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
		return errCode.NewClientError(errCode.CodeUserPwNotEqual)
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

func SignIn(psi *api.ParamSignIn) error {
	// 根据用户名查询加密密码
	password, err := app.GetUserPwByName(psi.Username)
	if err != nil {
		return err
	}
	// 比较数据库存储加密的密码和用户输入的密码
	compare := utils.BcryptCompare(password, psi.Password)
	if !compare {
		return errCode.NewClientError(errCode.CodeUserORPasswordErr)
	}
	// 生成jwt Token
	return nil
}
