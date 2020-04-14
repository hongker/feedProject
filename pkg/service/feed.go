package service

import (
	"encoding/json"
	"feedProject/pkg/app"
	"feedProject/pkg/constant"
	"feedProject/pkg/dao"
	"feedProject/pkg/entity"
	"feedProject/pkg/enum"
	"feedProject/pkg/request"
	"feedProject/pkg/response"
	"feedProject/pkg/utils"
	"fmt"
)

// 读扩散模式
type FeedService struct {

}

func Feed() *FeedService {
	return &FeedService{}
}

// Create
func (service *FeedService) Create(req request.FeedCreateRequest) error {
	feed := &entity.Feed{
		UserId:  req.UserId,
		CreatorId: req.UserId,
		Content: req.Content,
		FeedId:  int(utils.GenerateFeedId()),
		Status: enum.FeedStatusEnable,
	}

	if err := dao.Feed(app.DB()).Create(feed); err != nil {
		return fmt.Errorf("create:%v", err)
	}

	service.StoreCache(feed)

	return nil
}

// StoreCache 存储
func (service *FeedService) StoreCache(feed *entity.Feed) {
	// 更新缓存
	redisClient := app.Redis()
	cacheKey := constant.GetFeedInfoCacheKey(constant.GetFeedField(feed.CreatorId, feed.FeedId))
	item := response.FeedResponse{
		UserId: feed.UserId,
		Content:   feed.Content,
		CreatorId: feed.CreatorId,
		FeedId:    feed.FeedId,
		CreatedAt: feed.CreatedAt.String(),
	}

	res, _ := json.Marshal(item)
	redisClient.Set(cacheKey, res, constant.ExpireTime)
}

