package httpresponse

// 公共模块相关异常码
const (
	CODE_SUCCESS = 200 // 成功

	CODE_ERROR_SERVER = 500 // 服务器异常
	CODE_BAD_REQUEST  = 400 // 参数错误
	CODE_UNAUTHORIZED = 401 // 身份信息已过期
	CODE_FORBIDDEN    = 403 // 禁止访问
	CODE_NOT_FOUND    = 404 // 没找到内容
)

func (h *CommonResponseHandler) CodeText(code int) string {
	switch code {
	case CODE_SUCCESS:
		return "成功"
	case CODE_ERROR_SERVER:
		return "系统错误"
	case CODE_BAD_REQUEST:
		return "参数错误"
	case CODE_UNAUTHORIZED:
		return "身份信息已过期"
	case CODE_FORBIDDEN:
		return "禁止访问"
	case CODE_NOT_FOUND:
		return "没找到内容"
	default:
		return ""
	}
}
