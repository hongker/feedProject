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
		user.GET("/feed", handler.PullFeedHandler)
	}

	// relation api group
	relation := router.Group("v1/relation")
	{
		// 关注
		relation.POST("", handler.RelationCreateHandler)

		// 取消
		relation.DELETE("/:id", handler.RelationCancelHandler)
	}

	// feed api group
	feed := router.Group("v1/feed")
	{
		// 拉取feed
		feed.GET("", handler.FeedListHandler)

		// 创建
		feed.POST("", handler.FeedCreateHandler)

		// 更新
		feed.PUT("/:id", handler.FeedUpdateHandler)

		// 禁用
		feed.DELETE("/:id", handler.FeedDisableHandler)
	}

}
