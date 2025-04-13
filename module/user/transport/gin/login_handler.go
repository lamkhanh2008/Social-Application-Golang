package ginuser

import (
	"social_todo/common"
	"social_todo/component/tokenprovider/jwt"
	biz "social_todo/module/user/business"
	"social_todo/module/user/model"
	"social_todo/module/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Logic(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider("jwt", "200Lab.io")

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		business := biz.NewLogicBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		common.SuccessResponse(c, common.SimpleResponse(account))
	}
}
