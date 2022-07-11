package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"net/http"
)

type ResponseData struct {
	Code errCode.ErrorCode `json:"code"`
	Msg  string            `json:"msg"`
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
		code = errCode.ServerUnknown
		msg = err.Error()
	}

	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code errCode.ErrorCode, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: errCode.Success,
		Msg:  errCode.Success.GetMsg(),
		Data: data,
	})
}
