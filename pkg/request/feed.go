package request

type FeedCreateRequest struct {
	UserId int `json:"user_id"`
	Content string `json:"content"`
}


type FeedSearchRequest struct {
	Limit int `json:"limit"`
	CreatorIds []int `json:"creator_ids"`
	LastPullId int `json:"last_pull_id"`
}