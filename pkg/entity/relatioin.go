package entity

const(
	TableRelation = "relations"
)
// Relation 用户关注
type Relation struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	UserId int `json:"user_id" gorm:"column:user_id"`
	TargetId int `json:"target_id" gorm:"column:target_id"`
	Status int `json:"status" gorm:"column:status"`
	// 对time进行格式化
	CreatedAt Timestamp `gorm:"column:created_at" json:"created_at"`
	UpdatedAt Timestamp `gorm:"column:created_at" json:"updated_at"`
}

// TableName 指定模型的表名称
func (Relation) TableName() string {
	return TableRelation
}
