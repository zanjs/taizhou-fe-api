package response

type errorMessage struct {
	SUCCESS      string
	ERROR        string
	AuthErr      string
	RoleErr      string
	NotFound     string
	LoginError   string
	LoginTimeout string
	InActive     string
}

// ErrorMessage 错误码
var ErrorMessage = errorMessage{
	SUCCESS:      "成功",
	ERROR:        "错误",
	AuthErr:      "认证失败,请重新登陆",
	RoleErr:      "权限不够", // 权限不够
	NotFound:     "未找到",
	LoginError:   "1000", //用户名或密码错误
	LoginTimeout: "1001", //登录超时
	InActive:     "1002", //未激活账号
}
