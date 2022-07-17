package service

import (
	"github.com/lemonlzy/vegetableBlog/api"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/middleware"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/snowflake"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/utils"
)

// SignUp 注册
func SignUp(psu *api.ParamSignUp) error {
	if psu.Password != psu.RePassword {
		return errCode.NewClientError(errCode.UserPwNotEqual)
	}

	// 判断用户名是否重复
	if ok, err := app.CheckUserByName(psu.Username); ok || err != nil {
		return err
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

// SignIn 登录
func SignIn(psi *api.ParamSignIn) (*app.User, error) {
	// 根据用户名查询加密密码
	userBySQL, err := app.GetUserByName(psi.Username)
	if err != nil {
		return nil, err
	}

	// 比较数据库存储加密的密码和用户输入的密码
	equal := utils.BcryptCompare(userBySQL.Password, psi.Password)
	if !equal {
		return nil, errCode.NewClientError(errCode.UserORPasswordErr)
	}

	// 生成JWT token
	aToken, rToken, err := middleware.GenToken(userBySQL.UserID, psi.Username)
	if err != nil {
		return nil, err
	}

	// 组装响应的user结构体
	user := &app.User{
		UserID:   userBySQL.UserID,
		Nickname: userBySQL.Nickname,
		IsAdmin:  userBySQL.IsAdmin,
		AToken:   aToken,
		RToken:   rToken,
	}
	return user, nil
}

// ModifyPw 更改密码
func ModifyPw(pmp *api.ParamModifyPw) error {
	if pmp.RePassword != pmp.RePassword2 {
		return errCode.NewClientError(errCode.UserPwNotEqual)
	}

	// 判断用户是否存在
	if ok, err := app.CheckUserByID(pmp.UserID); !ok || err != nil {
		return err
	}

	// 校验原始密码
	pwBySQL, err := app.GetUserPwByID(pmp.UserID)
	if err != nil {
		return err
	}

	equal := utils.BcryptCompare(pwBySQL, pmp.Password)
	if !equal {
		return errCode.NewClientError(errCode.UserORPasswordErr)
	}

	// 密码加密
	hashPw := utils.BcryptPw(pmp.RePassword)
	user := &app.User{Password: hashPw}

	err = app.UpdateUserByID(pmp.UserID, user)
	if err != nil {
		return err
	}

	return nil
}
