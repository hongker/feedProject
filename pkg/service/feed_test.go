package service

import (
	"feedProject/pkg/request"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFeedService_Create(t *testing.T) {
	for i:=1;i<10;i++ {
		err := Feed().Create(request.FeedCreateRequest{
			UserId:  1,
			Content: fmt.Sprintf("bigVcontent:%d", i),
		})
		fmt.Println(err)
		assert.Nil(t, Feed().SyncQueue())
		time.Sleep(time.Second * 2)
	}
}

func TestFeedService_SyncQueue(t *testing.T) {
	assert.Nil(t, Feed().SyncQueue())
}


