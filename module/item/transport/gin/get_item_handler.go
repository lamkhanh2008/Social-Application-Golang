package ginitem

import (
	"fmt"
	"net/http"
	"social_todo/common"
	"social_todo/module/item/biz"
	"social_todo/module/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItemByID(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))
		fmt.Println(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewItemStorage(db)
		business := biz.NewItemBusiness(store)

		data, err := business.GetItemById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleResponse(data))
	}
}
