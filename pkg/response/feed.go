package response

type FeedResponse struct {
	UserId int `json:"user_id"`
	Content string `json:"content"`
	CreatorId int `json:"creator_id"`
	FeedId int `json:"feed_id"`
	CreatedAt string `json:"created_at"`
}
