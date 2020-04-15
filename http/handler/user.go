package handler

import (
	"feedProject/pkg/constant"
	"feedProject/pkg/enum"
	"feedProject/pkg/request"
	"feedProject/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PullNewFeedHandler 拉取新的feed信息
func PullNewFeedHandler(ctx *gin.Context)  {
	req := request.PullFeedRequest{
		Limit: constant.Limit,
	}
	// TODO 校验参数
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidParam,
			"message" : err.Error(),
		})
		return
	}

	items, err := service.User().GetNewFeed(req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidResponse,
			"message" : err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": enum.Success,
		"message": "success",
		"items": items,
		"offset": req.Offset + int64(len(items)),
	})
}

func PullHistoryFeedHandler(ctx *gin.Context)  {
	req := request.QueryHistoryFeedRequest{
		Limit: constant.Limit,
	}
	// TODO 校验参数
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidParam,
			"message" : err.Error(),
		})
		return
	}

	items, err := service.User().GetHistoryFeed(req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : enum.InvalidResponse,
			"message" : err.Error(),
		})
	}

	offset := req.Offset - int64(len(items))
	if offset < 0 {
		offset = 0
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": enum.Success,
		"message": "success",
		"items": items,
		"offset": offset,
	})
}
