package e

const (
	// 通用错误码
	SUCCESS              = 200
	ERROR                = 500
	INVALID_PARAMS       = 400
	ERROR_NOT_LOGIN      = 300
	ERROR_CAN_NOT_UPDATE = 30001
	ERROR_CAN_NOT_DELETE = 30002

	// 登录用户模块错误码
	ERROR_NOT_EXIST_ADMIN = 10001
	ERROR_PASSWORD        = 10002
	ERROR_FORBIDDEN_ADMIN = 10003

	// 广告模块
	ERROR_EXIST_AD     = 10004
	ERROR_NOT_EXIST_AD = 10005

	// 分栏模块
	ERROR_EXIST_CLASS            = 10006
	ERROR_NOT_EXIST_CLASS        = 10007
	ERROR_EXIST_ARTICLE_OF_CLASS = 10008

	ERROR_EXIST_TAG         = 40001
	ERROR_NOT_EXIST_TAG     = 40002
	ERROR_NOT_EXIST_ARTICLE = 40003

	ERROR_EXIST_ADMIN = 40005

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

func StatusText(code int) string {
	switch code {
	case SUCCESS:
		return "success"
	case ERROR:
		return "error"
	default:
		return ""
	}
}
