package ginitem

import (
	"fmt"
	"net/http"
	"social_todo/common"
	"social_todo/module/item/biz"
	"social_todo/module/item/model"
	"social_todo/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var itemData model.TodoItemCreation
		if err := ctx.ShouldBind(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println("asdasdads %+v", itemData)
		storage := storage.NewItemStorage(db)
		bussiness := biz.NewItemBusiness(storage)
		service := NewItemService(bussiness)
		err := service.biz.CreateItem(ctx.Request.Context(), &itemData)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusAccepted, common.SimpleResponse(itemData.Id))
	}
}
