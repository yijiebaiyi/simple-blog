package e

var MsgFlags = map[int]string{
	// 通用错误
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "请求参数错误",
	ERROR_NOT_LOGIN:      "请先登录",
	ERROR_CAN_NOT_UPDATE: "不允许修改",
	ERROR_CAN_NOT_DELETE: "不允许删除",

	// 登录用户模块错误
	ERROR_NOT_EXIST_ADMIN: "用户不存在",
	ERROR_EXIST_ADMIN:     "用户已存在",
	ERROR_PASSWORD:        "密码错误",
	ERROR_FORBIDDEN_ADMIN: "用户被禁用",

	// 广告模块
	ERROR_EXIST_AD: "广告已存在",
	ERROR_NOT_EXIST_AD: "广告不存在",

	// 分栏模块
	ERROR_EXIST_CLASS: "分栏已存在",
	ERROR_EXIST_ARTICLE_OF_CLASS: "分栏下存在文章",
	ERROR_NOT_EXIST_CLASS: "分栏不存在",

	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
