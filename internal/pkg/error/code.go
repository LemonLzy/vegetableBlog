package errCode

// 错误来源
const (
	SourceUnknown string = "Unknown"
	SourceClient  string = "Client"
	SourceServer  string = "Server"
)

// 通用返回码
const (
	Success            ErrorCode = iota
	ServerUnknown                = -1000 - iota // 未知错误
	ServerInvalidToken                          // 无效的token
	ClientReqInvalid                            // 请求参数错误
)

// User错误码
const (
	UserInvalidName   ErrorCode = -2000 - iota // 用户名格式不符合要求
	UserInvalidPass                            // 密码格式不符合要求
	UserORPasswordErr                          // 用户名或密码错误
	UserExist                                  // 用户名已存在
	UserNotExist                               // 用户不存在
	UserPwNotEqual                             // 用户两次输入的密码不相等
)

// Article错误码
const (
	ArticleCreate  ErrorCode = -3000 - iota // 新建Article错误
	ArticleUpdate                           // 更新Article错误
	ArticleDelete                           // 删除Article错误
	ArticleInvalid                          // Article不存在
)

// Tag错误码
const (
	TagCreate  ErrorCode = -4000 - iota // 新建Tag错误
	TagUpdate                           // 更新Tag错误
	TagDelete                           // 删除Tag错误
	TagInvalid                          // Tag不存在
)

var code2msg = map[ErrorCode]string{
	Success: "成功",

	ServerUnknown:      "未知错误",
	ServerInvalidToken: "无效的token",
	ClientReqInvalid:   "请求参数错误",

	UserInvalidName:   "用户名格式不符合要求",
	UserInvalidPass:   "密码格式不符合要求",
	UserORPasswordErr: "用户名或密码错误",
	UserExist:         "用户名已存在",
	UserNotExist:      "用户不存在",
	UserPwNotEqual:    "两次输入的密码不相等",

	ArticleCreate:  "新建Article错误",
	ArticleUpdate:  "更新Article错误",
	ArticleDelete:  "删除Article错误",
	ArticleInvalid: "Article不存在",

	TagCreate:  "新建Tag错误",
	TagUpdate:  "更新Tag错误",
	TagDelete:  "删除Tag错误",
	TagInvalid: "Tag不存在",
}

func (ei ErrorCode) GetMsg() string {
	msg, ok := code2msg[ei]
	if !ok {
		msg = code2msg[ServerUnknown]
	}
	return msg
}
