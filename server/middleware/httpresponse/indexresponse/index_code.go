package indexresponse

// 前台相关异常码, 2000x
const (
	CODE_ARTICLE_NOT_EXISTS      = 20001
	CODE_ARTICLE_EXISTS          = 20002
	CODE_COMMENTS_OVER_LIMIT_ONE = 20003
	CODE_COMMENTS_OVER_LIMIT_ALL = 20004
	CODE_SEARCH_NOT_FIND         = 20005
)

func (h *IndexResponseHandler) CodeText(code int) string {
	switch code {
	case CODE_ARTICLE_NOT_EXISTS:
		return "博文不存在😅"
	case CODE_ARTICLE_EXISTS:
		return "博文已存在😅"
	case CODE_COMMENTS_OVER_LIMIT_ONE:
		return "单条评论超过限制😣"
	case CODE_COMMENTS_OVER_LIMIT_ALL:
		return "所有评论超过限制😣"
	case CODE_SEARCH_NOT_FIND:
		return "搜索内容未找到，请换个关键词😅"
	default:
		return ""
	}
}
