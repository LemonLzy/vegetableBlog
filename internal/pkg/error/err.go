package errCode

import (
	"errors"
	"fmt"
)

type ErrorCode int16

type ErrorInfo struct {
	Source string
	Code   ErrorCode
	Msg    string
}

func NewError(Source string, code ErrorCode, msg string) *ErrorInfo {
	return &ErrorInfo{Source: Source, Code: code, Msg: msg}
}

func (ei *ErrorInfo) Error() string {
	return fmt.Sprintf("错误来源: %s, 错误码: %d, 错误信息: %s", ei.Source, ei.Code, ei.Msg)
}

func (ei *ErrorInfo) ErrSource() string {
	return ei.Source
}

func (ei *ErrorInfo) ErrCode() ErrorCode {
	return ei.Code
}

func (ei *ErrorInfo) ErrMsg() string {
	return ei.Msg
}

func NewClientError(code ErrorCode) *ErrorInfo {
	return NewError(SourceClient, code, code.GetMsg())
}

func NewClientErrorWithMsg(code ErrorCode, msg string) *ErrorInfo {
	return NewError(SourceClient, code, msg)
}

func NewServerError(code ErrorCode) *ErrorInfo {
	return NewError(SourceServer, code, code.GetMsg())
}

func NewUnknownError(code ErrorCode, msg string) *ErrorInfo {
	return NewError(SourceUnknown, code, msg)
}

func Err2ErrorInfo(err error) *ErrorInfo {
	var errType *ErrorInfo
	if errors.As(err, &errType) {
		return errType
	}
	return nil
}
