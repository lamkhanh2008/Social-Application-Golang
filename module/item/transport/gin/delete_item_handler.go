package gin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {

		}

	}
}
