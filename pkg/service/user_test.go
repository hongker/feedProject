package service

import (
	"feedProject/pkg/app"
	"feedProject/pkg/constant"
	"feedProject/pkg/request"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestUserService_PullFeed(t *testing.T) {
	req := request.PullFeedRequest{
		UserId: 2,
		Limit:  constant.Limit,
	}
	result, err := User().PullNewFeed(req)
	fmt.Println(result, err)
}

func TestUserService_PullHistoryFeed(t *testing.T) {
	req := request.QueryHistoryFeedRequest{
		UserId: 2,
		Limit:  constant.Limit,
		Offset:20,
	}

	result, err := User().PullHistoryFeed(req)
	assert.Nil(t, err)
	for _, item := range result {
		fmt.Println("id:", item.FeedId, "content", item.Content)
	}
}