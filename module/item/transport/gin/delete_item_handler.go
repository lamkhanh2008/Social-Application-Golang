package ginitem

import (
	"net/http"
	"social_todo/common"
	"social_todo/module/item/biz"
	"social_todo/module/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewItemStorage(db)
		business := biz.NewItemBusiness(store)

		err = business.DeleteItemById(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleResponse(true))

	}
}
