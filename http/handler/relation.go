package handler

import (
	"feedProject/pkg/enum"
	"feedProject/pkg/request"
	"feedProject/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RelationCreateHandler 关注
func RelationCreateHandler(ctx *gin.Context) {
	params := request.RelationRequest{}
	// TODO 校验参数

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidParam,
			"message" : err.Error(),
		})
		return
	}

	if err := service.Relation().Create(params.UserId, params.TargetId); err != nil {
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

// RelationCancelHandler 取消
func RelationCancelHandler(ctx *gin.Context)  {
	params := request.RelationRequest{}
	// TODO 校验参数

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidParam,
			"message" : err.Error(),
		})
		return
	}

	if err := service.Relation().Cancel(params.UserId, params.TargetId); err != nil {
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
