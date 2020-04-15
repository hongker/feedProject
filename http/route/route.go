package route

import (
	"feedProject/http/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(router *gin.Engine)  {
	// home page
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "feed project")
	})

	// user api group
	user := router.Group("v1/user")
	{
		// 获取feed流信息
		user.GET("feed/new", handler.PullNewFeedHandler)

		// 获取feed流信息
		user.GET("feed/history", handler.PullHistoryFeedHandler)
	}


}
