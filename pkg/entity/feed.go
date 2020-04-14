package entity

const(
	TableFeed = "feed"
)
// Feed 信息流
type Feed struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	UserId int `json:"user_id" gorm:"column:user_id"`
	CreatorId int `json:"creator_id" gorm:"column:creator_id"`
	FeedId int `json:"feed_id"`
	Content string `json:"content"`
	Status int `json:"status" gorm:"column:status"`
	// 对time进行格式化
	CreatedAt Timestamp `gorm:"column:created_at" json:"created_at"`
	UpdatedAt Timestamp `gorm:"column:created_at" json:"updated_at"`
}

// TableName 指定模型的表名称
func (Feed) TableName() string {
	return TableFeed
}
