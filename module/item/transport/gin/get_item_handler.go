package gin

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetItemByID(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
	}
}
