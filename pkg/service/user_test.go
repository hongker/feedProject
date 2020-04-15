package service

import (
	"feedProject/pkg/app"
	"feedProject/pkg/constant"
	"feedProject/pkg/request"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMain(m *testing.M)  {
	app.ConnectDB()
	app.ConnectRedis()

	m.Run()
}

func TestUserService_Create(t *testing.T) {
	for i:=1;i<100 ; i++ {
		err := User().Create(fmt.Sprintf("test:%d", i))

		assert.Nil(t, err)
	}

}

func TestUserService_GetNewFeed(t *testing.T) {
	req := request.PullFeedRequest{
		UserId: 2,
		Limit:  5,
		Offset: 9,
	}

	result, err := User().GetNewFeed(req)
	assert.Nil(t, err)
	for _, item := range result {
		fmt.Printf("creted_at:%s,content:%s\n", item.CreatedAt, item.Content)
	}

}

func TestUserService_GetHistoryFeed(t *testing.T) {
	req := request.QueryHistoryFeedRequest{
		UserId: 2,
		Limit:  constant.Limit,
		Offset: 5,
	}

	result, err := User().GetHistoryFeed(req)
	assert.Nil(t, err)
	for _, item := range result {
		fmt.Printf("creted_at:%s,content:%s\n", item.CreatedAt, item.Content)
	}
}

func TestUserService_InitFeedOfInactive(t *testing.T) {
	assert.Nil(t, User().InitFeedOfInactive(2))
}
func TestZRange(t *testing.T)  {
	fmt.Println(time.Now().Nanosecond() / 1000)
}