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
	"sync"
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

// PullFeed 拉取最新的feed
func (service *UserService) PullNewFeed(req request.PullFeedRequest) ([]response.FeedResponse, error) {
	var items []response.FeedResponse

	// feed fields 数组
	var fields []string
	redisClient := app.Redis()

	//
	ctx, cancel := context.WithCancel(context.Background())

	//  获取指定数量关注的用户
	userRelationOffsetKey := constant.GetUserRelationOffsetCacheKey(req.UserId)
	lastPullIdKey := constant.GetUserPullNewFeedLastId(req.UserId)

	// 默认成功
	offset, _ := redisClient.Get(userRelationOffsetKey).Int()
	lastPullId, _ := redisClient.Get(lastPullIdKey).Int()
	relationDao := dao.Relation(app.DB())
	ch := make(chan string, constant.UserRelationLimit)

	// 使用协程
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("complete")
				// 关闭通道
				close(ch)
				return
			default:
				fmt.Println("working...")
				creatorIds, err := relationDao.GetTargetIdByUserIdAndOffset(req.UserId, offset, constant.UserRelationLimit)
				if err != nil || len(creatorIds) == 0{
					close(ch)
					// 获取数据失败，退出协程
					return
				}

				// 更新offset,用于下次循环
				offset += len(creatorIds)
				// 获取这些用户最新的feed
				// TODO 可以优化为读取缓存
				feeds , err := dao.Feed(app.DB()).SearchFeed(request.FeedSearchRequest{
					Limit:      req.Limit,
					CreatorIds: creatorIds,
					LastPullId: lastPullId,
				})

				if err != nil || len(feeds) == 0 {
					continue
				}
				lastPullId = feeds[len(feeds)-1].ID
				redisClient.Set(lastPullIdKey, lastPullId, constant.ExpireTime)

				for _, feed := range feeds {
					ch <- constant.GetFeedField(feed.CreatorId, feed.FeedId)
				}

				if len(feeds) >= req.Limit {
					close(ch)
					return
				}
			}
		}
	}()

	mx := new(sync.Mutex)

	for field := range ch {
		fmt.Println(field)
		if len(fields) >= req.Limit {
			mx.Lock()
			cancel()
			mx.Unlock()
			break
		}
		fields = append(fields, field)
	}

	items = service.getFormatFeedResponseItems(fields)

	// 异步写入history队列
	service.WriteHistoryFeedAsync(req.UserId, items)


	return items, nil
}

// WriteHistoryFeedAsync 异步写入历史
func (service *UserService) WriteHistoryFeedAsync(userId int, items []response.FeedResponse) error {
	// 使用协程模拟异步事件同步历史feed
	redisClient := app.Redis()
	cacheKey := constant.GetUserHistoryFeedListCacheKey(userId)
	for _, item := range items {
		_ = redisClient.LPush(cacheKey, constant.GetFeedField(item.CreatorId, item.FeedId))
	}

	return nil
}

// PullHistoryFeed 拉取历史feed
func (service *UserService) PullHistoryFeed(req request.QueryHistoryFeedRequest) ([]response.FeedResponse, error) {
	var items []response.FeedResponse
	// 查询历史feed列表
	cacheKey := constant.GetUserHistoryFeedListCacheKey(req.UserId)
	redisClient := app.Redis()

	// 从左往右读取,默认左边是最新的数据,右边是老旧数据
	result, err := redisClient.LRange(cacheKey, req.Offset, req.Offset+req.Limit).Result()
	if err != nil {
		return nil, fmt.Errorf("get list:%v", err)
	}

	items = service.getFormatFeedResponseItems(result)

	// 判断数量是否大于请求需要的数量
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
