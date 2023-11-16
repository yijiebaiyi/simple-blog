package indexresponse

import (
	"gin_admin/middleware/httpresponse"

	"github.com/gin-gonic/gin"
)

type IndexResponseHandler struct {
}

func IndexResponse(ctx *gin.Context, code int, err error, data ...interface{}) {
	response := httpresponse.CommonResponse{
		Code: code,
		Err:  err,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}

	httpresponse.SetResponseHandler(&IndexResponseHandler{}, &response)
	httpresponse.JsonResponse(ctx, response)
}
