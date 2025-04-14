package ginitem

import (
	"net/http"
	"social_todo/common"
	"social_todo/module/item/biz"
	"social_todo/module/item/model"
	"social_todo/module/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItemById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		var dataUpdate model.TodoItemUpdate
		err = ctx.ShouldBind(&dataUpdate)
		if err != nil {
			panic(err)
		}

		requester := ctx.MustGet(common.CurrrentUser).(common.Requester)
		store := storage.NewItemStorage(db)
		business := biz.NewUpdateItemBusiness(store, requester)

		err = business.UpdateItem(ctx.Request.Context(), id, &dataUpdate)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleResponse(true))

	}
}
