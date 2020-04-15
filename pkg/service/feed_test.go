package service

import (
	"feedProject/pkg/request"
	"fmt"
	"testing"
	"time"
)

func TestFeedService_Create(t *testing.T) {
	for i:=1;i<=100;i++ {
		err := Feed().Create(request.FeedCreateRequest{
			UserId:  3,
			Content: fmt.Sprintf("commonContent:%d", i),
		})
		fmt.Println(err)
		time.Sleep(time.Second * 1)
	}
}

func TestFeedService_SyncQueue(t *testing.T) {
	for {
		if err := Feed().SyncQueue(); err != nil {
			break
		}
	}
}


