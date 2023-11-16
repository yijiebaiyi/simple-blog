package httpresponse

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResponseHandler struct {
}

type MessageProvider interface {
	CodeText(int) string
}

type CommonResponse struct {
	Code            int         `json:"code"`
	Message         string      `json:"message"`
	Data            interface{} `json:"data"`
	Err             error       `json:"-"`
	MessageProvider `json:"-"`
}

func SetResponseHandler(provider MessageProvider, response *CommonResponse) {
	response.MessageProvider = provider
}

func JsonResponse(ctx *gin.Context, response CommonResponse) {
	if response.MessageProvider == nil {
		response.MessageProvider = &CommonResponseHandler{}
	}
	if response.Message == "" {
		response.Message = response.CodeText(response.Code)
	}
	if response.Err != nil {
		fmt.Printf("%v", response.Err)
		// logger.Logger.Info("接口请求返回失败，【返回信息】%v, 【错误信息】%v", response, response.Err.Error())
	}
	ctx.JSON(http.StatusOK, response)
}

func SuccessResponse(ctx *gin.Context, data ...interface{}) {
	var response CommonResponse
	response.Code = CODE_SUCCESS
	if len(data) > 0 {
		response.Data = data[0]
	}
	JsonResponse(ctx, response)
}

func ErrorResponse(ctx *gin.Context, code int, err error, data ...interface{}) {
	var response CommonResponse
	response.Code = code
	response.Err = err
	if len(data) > 0 {
		response.Data = data[0]
	}
	JsonResponse(ctx, response)
}
