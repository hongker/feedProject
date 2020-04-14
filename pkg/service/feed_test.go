package service

import (
	"feedProject/pkg/request"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeedService_Create(t *testing.T) {
	for i:=1; i<100;i++ {
		req := request.FeedCreateRequest{
			UserId:  1,
			Content: fmt.Sprintf("content:%d", i),
		}
		err := Feed().Create(req)
		assert.Nil(t, err)
	}
}


