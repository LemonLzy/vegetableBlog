package error

// 通用错误码
const (
	Success                   = iota         // 请求成功
	ErrServerAuthInvalidToken = -1000 - iota // 无效的token
	ErrServerUnknown                         // 未知错误
)

// User错误码
const (
	ErrUserInvalid         = -2000 - iota // 用户名格式不符合要求
	ErrUserPasswordInvalid                // 密码格式不符合要求
	ErrUserORPassword                     // 用户名或密码错误
)

// Article错误码
const (
	ErrBlogCreate = -3000 - iota // 密码错误
	ErrBlogUpdate                // 服务端参数验证失败
	ErrBlogDelete
	ErrBlogNotExist
)
