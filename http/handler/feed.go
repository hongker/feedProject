package handler

import (
	"feedProject/pkg/enum"
	"feedProject/pkg/request"
	"feedProject/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FeedListHandler(ctx *gin.Context)  {

}

func FeedCreateHandler(ctx *gin.Context)  {
	req := request.FeedCreateRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidParam,
			"message" : err.Error(),
		})
		return
	}

	if err := service.Feed().Create(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidResponse,
			"message" : err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": enum.Success,
		"message": "success",
	})
}

func FeedUpdateHandler(ctx *gin.Context)  {

}

func FeedDisableHandler(ctx *gin.Context)  {

}
