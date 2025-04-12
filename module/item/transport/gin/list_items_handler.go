package ginitem

import (
	"net/http"
	"social_todo/common"
	"social_todo/module/item/biz"
	"social_todo/module/item/model"
	"social_todo/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListItems(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var querytring struct {
			common.Paging
			model.Filter
		}

		if err := ctx.ShouldBind(&querytring); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		querytring.Paging.Process()

		store := storage.NewItemStorage(db)
		business := biz.NewItemBusiness(store)

		result, err := business.ListItems(ctx.Request.Context(), &querytring.Filter, &querytring.Paging)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(result, querytring.Filter, querytring.Paging))
	}
}
