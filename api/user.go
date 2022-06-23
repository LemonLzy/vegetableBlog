package api

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username,omitempty" binding:"required"`
	Password   string `json:"password,omitempty" binding:"required"`
	RePassword string `json:"re_password,omitempty" binding:"required"`
}

// ParamSignIn 登录请求参数
type ParamSignIn struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
