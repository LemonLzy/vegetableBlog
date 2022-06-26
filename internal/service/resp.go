package service

import (
	"github.com/gin-gonic/gin"
	errCode "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"net/http"
)

type ResponseData struct {
	Code errCode.ErrorCode `json:"code"`
	Msg  interface{}       `json:"msg"`
	Data interface{}       `json:"data"`
}

func ResponseError(c *gin.Context, err error) {
	var code errCode.ErrorCode
	var msg string

	if _, ok := err.(*errCode.ErrorInfo); ok {
		errType := errCode.Err2ErrorInfo(err)
		code = errType.Code
		msg = code.GetMsg()
	} else {
		code = errCode.CodeServerUnknown
		msg = err.Error()
	}

	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code errCode.ErrorCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: errCode.CodeSuccess,
		Msg:  errCode.CodeSuccess.GetMsg(),
		Data: data,
	})
}
