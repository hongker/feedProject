package request

type PullFeedRequest struct {
	UserId int `form:"user_id"`
	Limit int `form:"limit"`

}

// QueryHistoryFeedRequest 查询历史feed
type QueryHistoryFeedRequest struct {
	UserId int `form:"user_id"`
	Offset int64 `form:"offset"`
	Limit int64 `form:"limit"`
}
