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

// ParamModifyPw 更改密码请求参数
type ParamModifyPw struct {
	UserID      int64  `json:"user_id,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
	RePassword  string `json:"re_password,omitempty" binding:"required"`
	RePassword2 string `json:"re_password2,omitempty" binding:"required"`
}

// ParamDel 删除用户请求参数
type ParamDel struct {
	UserID int64 `json:"user_id,omitempty" binding:"required"`
}
