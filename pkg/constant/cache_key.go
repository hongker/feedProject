package constant

import (
	"fmt"
)

const(
	historyFeedListPrefix = "cache_history_feed_list"
	userFeedInfoPrefix = "cache_user_feed_info"
	userFeedOffsetPrefix = "cache_user_feed_offset"
	userRelationOffsetPrefix = "cache_user_relation_offset"
	userNewFeedPrefix = "cache_user_new_feed"
	userPullLastIdPrefix = "cache_user_pull_last_id"
	userFeedQueuePrefix = "cache_feed_queue"

	FeedSyncQueue = "cache_feed_sync_queue"
)


// GetUserHistoryFeedListCacheKey 获取用户历史feed的list,实现分页
// key=history_feed_list:userId
// value=creator_uid:feedId
func GetUserHistoryFeedListCacheKey(userId int) string  {
	return fmt.Sprintf("%s:%d", historyFeedListPrefix, userId)
}

// GetFeedCacheKey 获取用户的feed内容key,通过hash存储的是feedJson
// key=user_field_info:creator_id:feed_id
// value= {"content":content,"created_at":created_at...}
func GetFeedInfoCacheKey(field string) string {
	return fmt.Sprintf("%s:%s", userFeedInfoPrefix, field)
}

// GetUserRelationOffsetCacheKey 获取用户关注offset
func GetUserRelationOffsetCacheKey(userId int) string {
	return fmt.Sprintf("%s:%d", userRelationOffsetPrefix, userId)
}

// GetUserFeedOffsetCacheKey 获取用户拉取feed的offset的key，存的是int
func GetUserFeedOffsetCacheKey(userId int) string {
	return fmt.Sprintf("%s:%d", userFeedOffsetPrefix, userId)
}

// GetUserPullNewFeedLastId 获取用户最后一条拉取新Feed的Id
func GetUserPullNewFeedLastId(userId int) string  {
	return fmt.Sprintf("%s:%d", userPullLastIdPrefix, userId)
}

// GetUserFeedQueueKey 获取用户的feed队列，类型为sorted set
// key=user_feed_queue:userId
// value={"content":content,"created_at":created_at..}
func GetUserFeedQueueKey(userId int) string {
	return fmt.Sprintf("%s:%d", userFeedQueuePrefix, userId)
}

// GetFeedField
func GetFeedField(creatorId, feedId int) string {
	return fmt.Sprintf("%d:%d", creatorId, feedId)
}

