package ginitem

import (
	"fmt"
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
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println("sss")
		var dataUpdate model.TodoItemUpdate
		err = ctx.ShouldBind(&dataUpdate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewItemStorage(db)
		business := biz.NewItemBusiness(store)

		err = business.UpdateItem(ctx.Request.Context(), id, &dataUpdate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleResponse(true))

	}
}
