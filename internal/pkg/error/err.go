package errCode

import "fmt"

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

func NewClientError(code ErrorCode, msg string) *ErrorInfo {
	return NewError(SourceClient, code, msg)
}

func NewServerError(code ErrorCode, msg string) *ErrorInfo {
	return NewError(SourceServer, code, msg)
}

func NewUnknownError(code ErrorCode, msg string) *ErrorInfo {
	return NewError(SourceUnknown, code, msg)
}
