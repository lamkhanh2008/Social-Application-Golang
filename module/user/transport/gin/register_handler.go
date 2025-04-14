package ginuser

import (
	"social_todo/common"
	biz "social_todo/module/user/business"
	"social_todo/module/user/model"
	"social_todo/module/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()
		biz := biz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		// data.Mask(common.DBTypeUser)

		common.SuccessResponse(c, common.SimpleResponse(data))
	}
}
