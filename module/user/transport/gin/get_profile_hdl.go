package ginuser

import (
	"net/http"
	"social_todo/common"

	"github.com/gin-gonic/gin"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrrentUser)

		c.JSON(http.StatusOK, common.SimpleResponse(u))
	}
}
