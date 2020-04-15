package utils

import (
	"time"
)

// GenerateFeedId 生成唯一的feedId,同时可以按时间进行排序
func GenerateFeedId() int64  {
	return time.Now().Unix() * 1000000 + int64(time.Now().Nanosecond() / 1000)
}
