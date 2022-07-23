package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/api"
	errCode "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/resp"
	"github.com/lemonlzy/vegetableBlog/internal/service"
	"github.com/lemonlzy/vegetableBlog/pkg"
	"strconv"
)

// UserSignUpHandler 用户注册
func UserSignUpHandler(c *gin.Context) {
	psu := new(api.ParamSignUp)
	err := c.ShouldBindJSON(psu)
	if err != nil {
		resp.ResponseError(c, errCode.NewClientError(errCode.ClientReqInvalid))
		return
	}

	err = service.SignUp(psu)
	if err != nil {
		resp.ResponseError(c, err)
		return
	}

	resp.ResponseSuccess(c, nil)
}

// UserSignInHandler 用户登录
func UserSignInHandler(c *gin.Context) {
	psi := new(api.ParamSignIn)
	err := c.ShouldBindJSON(psi)
	if err != nil {
		resp.ResponseError(c, errCode.NewClientError(errCode.ClientReqInvalid))
		return
	}

	user, err := service.SignIn(psi)
	if err != nil {
		resp.ResponseError(c, err)
		return
	}

	resp.ResponseSuccess(c, user)
}

// UserSignOutHandler 用户注销
func UserSignOutHandler(c *gin.Context) {
}

// UserModifyHandler 用户修改基本信息
func UserModifyHandler(c *gin.Context) {

}

// UserModifyPwHandler 用户修改密码
func UserModifyPwHandler(c *gin.Context) {
	pmp := new(api.ParamModifyPw)
	err := c.ShouldBindJSON(pmp)
	if err != nil {
		resp.ResponseError(c, errCode.NewClientError(errCode.ClientReqInvalid))
		return
	}

	err = service.ModifyPw(pmp)
	if err != nil {
		resp.ResponseError(c, err)
		return
	}

	resp.ResponseSuccess(c, nil)
}

// UserFontPageHandler 用户个人首页信息统计
func UserFontPageHandler(c *gin.Context) {
	// 获取用户id
	id := c.Param("userID")
	userID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		resp.ResponseError(c, errCode.NewClientError(errCode.ClientReqInvalid))
		return
	}

	page, size := pkg.GetPageInfo(c)
	// 文章信息
	articles, err := service.ArticleCount(userID, page, size)

	// tag信息
	tags, err := service.TagCount()

	// 用户基本信息

	// 信息组装
	data := make(map[string]interface{}, 4)
	data["articles"] = articles
	data["tags"] = tags
	resp.ResponseSuccess(c, data)
}
