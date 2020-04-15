package service

import (
	"context"
	"encoding/json"
	"feedProject/pkg/app"
	"feedProject/pkg/constant"
	"feedProject/pkg/dao"
	"feedProject/pkg/entity"
	"feedProject/pkg/request"
	"feedProject/pkg/response"
	"fmt"
	"github.com/go-redis/redis"
)

type UserService struct {
}

func User() *UserService {
	return &UserService{}
}

func (service *UserService) Create(username string) error {
	user := &entity.User{
		Username:username,
	}

	return dao.User(app.DB()).Create(user)
}

// GetNewFeed 获取新Feed数据
func (service *UserService) GetNewFeed(req request.PullFeedRequest) ([]response.FeedResponse, error) {
	var items []response.FeedResponse
	// 查询历史feed列表
	cacheKey := constant.GetUserFeedQueueKey(req.UserId)
	redisClient := app.Redis()

	// 通过redis的sorted set排序
	result, err := redisClient.ZRange(cacheKey, req.Offset, req.Offset+req.Limit-1).Result()
	if err != nil {
		return nil, fmt.Errorf("get list:%v", err)
	}

	items = service.getFormatFeedResponseItems(result)

	return items, nil
}

// GetHistoryFeed 拉取历史feed
func (service *UserService) GetHistoryFeed(req request.QueryHistoryFeedRequest) ([]response.FeedResponse, error) {
	if req.Offset == 0 {
		return nil, nil
	}
	var items []response.FeedResponse
	// 查询历史feed列表
	cacheKey := constant.GetUserFeedQueueKey(req.UserId)
	redisClient := app.Redis()

	offset := req.Offset - req.Limit
	if offset < 0 {
		offset = 0
	}


	// 通过redis的sorted set排序
	result, err := redisClient.ZRange(cacheKey, offset, req.Offset-1).Result()
	if err != nil {
		return nil, fmt.Errorf("get list:%v", err)
	}

	items = service.getFormatFeedResponseItems(result)

	return items, nil
}


// formatFeedResponseItems 格式化
// fields feed的key的数组
func (service *UserService) getFormatFeedResponseItems(fields []string) []response.FeedResponse {
	result := []response.FeedResponse{}

	// TODO 从hash中获取
	var feedKeys []string
	for _, field := range fields {
		feedKeys = append(feedKeys, constant.GetFeedInfoCacheKey(field))
	}

	redisClient := app.Redis()
	// 读缓存
	// TODO 默认缓存是高可用的，一定存在，如果出现异常，需采取补偿措施写入缓存数据,这里就不实现了
	existCacheItems, _ := redisClient.MGet(feedKeys...).Result()

	// 解析json
	for _, item := range existCacheItems {
		var resp response.FeedResponse
		_ = json.Unmarshal([]byte(item.(string)), &resp)

		result = append(result, resp)
	}

	return result
}

// InitFeedOfInactive 对于不活跃的用户，再次登录时初始化feed流
func (service *UserService) InitFeedOfInactive(userId int) error  {
	redisClient := app.Redis()

	// 超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 1e9 * 5)
	defer cancel()
	ch := make(chan entity.Feed, constant.UserRelationLimit)

	go func() {
		// 分页查询关注的人
		lastId := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("complete")
				// 关闭通道
				close(ch)
				return
			default:
				fmt.Println("working...")
				feeds, err := dao.Feed(app.DB()).GetFollowUserFeedItems(userId, lastId, constant.Limit)
				if err != nil || len(feeds) == 0 {
					close(ch)
					return
				}

				lastId = feeds[len(feeds)-1].ID
				for _, feed := range feeds {
					ch <- feed
				}
			}
		}
	}()

	// 使用Redis有序集合排序
	feedCacheKey := constant.GetUserFeedQueueKey(userId)
	for item := range ch {
		redisClient.ZAdd(feedCacheKey, redis.Z{
			Score:  float64(item.FeedId),
			Member: constant.GetFeedField(item.CreatorId, item.FeedId),
		})
	}

	return nil
}