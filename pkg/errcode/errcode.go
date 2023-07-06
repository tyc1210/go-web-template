package errcode

var (
	Success           = NewCode(0, "成功")
	ServerError       = NewCode(10000, "服务内部错误")
	InvalidParams     = NewCode(10001, "入参错误")
	UnauthorizedToken = NewCode(10002, "鉴权失败")

	UserSave  = NewCode(20001, "创建用户失败")
	UserLogin = NewCode(20002, "用户登录失败")

	FileError     = NewCode(30001, "文件上传失败")
	FileOverLimit = NewCode(30002, "文件过大")
	FileErrorExt  = NewCode(30003, "文件格式错误")
)

type Code struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

func NewCode(c int, s string) Code {
	return Code{code: c, msg: s}
}

func (code Code) SetMsg(msg string) {
	code.msg = msg
}

func (code Code) GetCode() int {
	return code.code
}

func (code Code) GetMsg() string {
	return code.msg
}
