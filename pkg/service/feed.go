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
	"github.com/go-redis/redis"
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

	service.pushQueue(feed.ID)

	return nil
}

// pushQueue 推送到队列里，异步消费
func (service *FeedService) pushQueue(id int) {
	app.Redis().LPush(constant.FeedSyncQueue, id)
}

// SyncQueue 从队列中同步缓存
func (service *FeedService) SyncQueue() error {
	// 更新缓存
	redisClient := app.Redis()

	id, err := redisClient.RPop(constant.FeedSyncQueue).Int()
	if err != nil {
		return err
	}

	feed, err := dao.Feed(app.DB()).Get(id)
	if err != nil {
		return err
	}

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

	// TODO 如果是大V(role_type=2),则往活跃用户主动推送feed信息
	user , err := dao.User(app.DB()).Get(feed.CreatorId)
	if err != nil {
		return err
	}

	if user.RoleType == enum.UserRoleTypeNormal {
		// 小博主直接推送,因为关注的人不多
		followUserIds, err := dao.Relation(app.DB()).GetFollowUserIds(feed.CreatorId)
		if err != nil {
			return err
		}
		_ = service.Push(followUserIds, feed)
	}else {
		// 大V只推送给活跃用户，活跃粉丝占全部粉丝的小部分
		livedFollowUserIds, err := dao.Relation(app.DB()).GetActiveFollowUserIds(feed.CreatorId)
		if err != nil {
			return err
		}

		_ = service.Push(livedFollowUserIds, feed)
	}
	return nil
}

// Push 写扩散(推模式)
func (service *FeedService) Push(userIds []int, feed *entity.Feed) error {
	redisClient := app.Redis()
	for _, userId := range userIds {
		feedQueueKey := constant.GetUserFeedQueueKey(userId)
		_ = redisClient.ZAdd(feedQueueKey, redis.Z{
			Score:  float64(feed.FeedId),
			Member: constant.GetFeedField(feed.CreatorId, feed.FeedId),
		})
	}
	return nil

}

