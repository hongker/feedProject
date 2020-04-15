package request

type PullFeedRequest struct {
	UserId int `form:"user_id"`
	Limit int64 `form:"limit"`
	Offset int64 `form:"offset"`

}

// QueryHistoryFeedRequest 查询历史feed
type QueryHistoryFeedRequest struct {
	UserId int `form:"user_id"`
	Offset int64 `form:"offset"`
	Limit int64 `form:"limit"`
}
