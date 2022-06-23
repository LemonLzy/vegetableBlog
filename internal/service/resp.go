package service

import (
	"github.com/gin-gonic/gin"
	err "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"net/http"
)

type ResponseData struct {
	Code err.ErrorCode `json:"code"`
	Msg  interface{}   `json:"msg"`
	Data interface{}   `json:"data"`
}

func ResponseError(c *gin.Context, code err.ErrorCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code err.ErrorCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: err.CodeSuccess,
		Msg:  err.CodeSuccess.GetMsg(),
		Data: data,
	})
}
