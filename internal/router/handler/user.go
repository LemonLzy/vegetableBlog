package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/api"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	errCode "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"github.com/lemonlzy/vegetableBlog/internal/service"
	"net/http"
	"strconv"
)

func UserCreateHandler(c *gin.Context) {
	u := new(app.User)

	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数绑定失败",
		})
		return
	}

	// 判断用户是否存在
	if ok, err := app.GetUserByName(u.Username); ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户名重复" + err.Error(),
		})
		return
	}

	if err := app.CreateUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}

func UserUpdateHandler(c *gin.Context) {
	// 获取用户id
	idStr := c.Param("id")
	userID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户ID获取失败",
		})
		return
	}

	u := new(app.User)
	if err = c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数绑定失败",
		})
		return
	}

	err = app.UpdateUserByID(userID, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新用户失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}

func UserSignUpHandler(c *gin.Context) {
	psu := new(api.ParamSignUp)
	err := c.ShouldBindJSON(psu)
	if err != nil {
		service.ResponseError(c, errCode.NewClientError(errCode.CodeClientReqInvalid))
		return
	}

	err = service.SignUp(psu)
	if err != nil {
		service.ResponseError(c, err)
		return
	}

	service.ResponseSuccess(c, nil)
}

func UserSignInHandler(c *gin.Context) {
	psi := new(api.ParamSignIn)
	err := c.ShouldBindJSON(psi)
	if err != nil {
		service.ResponseError(c, errCode.NewClientError(errCode.CodeClientReqInvalid))
		return
	}

	err = service.SignIn(psi)
	if err != nil {
		service.ResponseError(c, err)
		return
	}

	service.ResponseSuccess(c, nil)
}

func UserModifyHandler(c *gin.Context) {

}

func UserModifyPwHandler(c *gin.Context) {
	pmp := new(api.ParamModifyPw)
	err := c.ShouldBindJSON(pmp)
	if err != nil {
		service.ResponseError(c, errCode.NewClientError(errCode.CodeClientReqInvalid))
		return
	}

	err = service.ModifyPw(pmp)
	if err != nil {
		service.ResponseError(c, err)
		return
	}

	service.ResponseSuccess(c, nil)
}
