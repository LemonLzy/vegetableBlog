package errCode

const CodeSuccess ErrorCode = 0 // 请求成功

// 通用错误码
const (
	_                      = iota
	CodeServerUnknown      = -1000 - iota // 未知错误
	CodeServerInvalidToken                // 无效的token
	CodeClientReqInvalid                  // 请求参数错误
)

// User错误码
const (
	CodeUserInvalidName ErrorCode = -2000 - iota // 用户名格式不符合要求
	CodeUserInvalidPass                          // 密码格式不符合要求
	CodeUserORPassword                           // 用户名或密码错误
)

// Article错误码
const (
	CodeArticleCreate  ErrorCode = -3000 - iota // 新建Article错误
	CodeArticleUpdate                           // 更新Article错误
	CodeArticleDelete                           // 删除Article错误
	CodeArticleInvalid                          // Article不存在
)

// Tag错误码
const (
	CodeTagCreate  ErrorCode = -4000 - iota // 新建Tag错误
	CodeTagUpdate                           // 更新Tag错误
	CodeTagDelete                           // 删除Tag错误
	CodeTagInvalid                          // Tag不存在
)

var codeMsgMap = map[ErrorCode]string{
	CodeSuccess:            "CodeSuccess",
	CodeServerUnknown:      "未知错误",
	CodeServerInvalidToken: "无效的token",
	CodeClientReqInvalid:   "请求参数错误",

	CodeUserInvalidName: "用户名格式不符合要求",
	CodeUserInvalidPass: "密码格式不符合要求",
	CodeUserORPassword:  "用户名或密码错误",

	CodeArticleCreate:  "新建Article错误",
	CodeArticleUpdate:  "更新Article错误",
	CodeArticleDelete:  "删除Article错误",
	CodeArticleInvalid: "Article不存在",

	CodeTagCreate:  "新建Tag错误",
	CodeTagUpdate:  "更新Tag错误",
	CodeTagDelete:  "删除Tag错误",
	CodeTagInvalid: "Tag不存在",
}

const (
	SourceUnknown string = "Unknown"
	SourceClient  string = "Client"
	SourceServer  string = "Server"
)

func (ei ErrorCode) GetMsg() string {
	msg, ok := codeMsgMap[ei]
	if !ok {
		msg = codeMsgMap[CodeServerUnknown]
	}
	return msg
}
