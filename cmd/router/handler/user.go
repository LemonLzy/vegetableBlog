package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/api"
	"github.com/lemonlzy/vegetableBlog/internal/app"
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
	p := new(api.ParamSignUp)
	err := c.ShouldBindJSON(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "缺少必填参数",
		})
		return
	}
}

func UserSignInHandler(c *gin.Context) {

}
