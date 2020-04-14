package utils

import (
	"time"
)

func GenerateFeedId() int64  {
	return time.Now().Unix() * 10000 + int64(time.Now().Nanosecond())
}
