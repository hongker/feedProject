package request

type RelationRequest struct {
	UserId int `json:"user_id"`
	TargetId int `json:"target_id"`
}
