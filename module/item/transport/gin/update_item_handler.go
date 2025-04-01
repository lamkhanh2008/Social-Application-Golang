package gin

import (
	"social_todo/module/item/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateItemById(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {

		}

		var dataUpdate model.TodoItemUpdate
		err = ctx.ShouldBind(&dataUpdate)
	}
}
