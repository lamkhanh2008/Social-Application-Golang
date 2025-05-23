package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Extra  interface{} `json:"extra,omitempty"`
}

func NewResponse(data, paging, extra interface{}) *successResponse {
	return &successResponse{
		Data:   data,
		Paging: paging,
		Extra:  extra,
	}
}

func SimpleResponse(data interface{}) *successResponse {
	return &successResponse{
		Data: data,
	}
}

func SuccessResponse(ctx *gin.Context, data *successResponse) {
	ctx.JSON(http.StatusOK, data)
}
