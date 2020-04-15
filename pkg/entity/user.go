package entity

const (
	TableUser = "users"
)

// User 用户
type User struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Username string `json:"username" gorm:"column:username"`
	RoleType int `json:"role_type" gorm:"column:role_type"`
	Status int `json:"status" gorm:"column:status"`
	// 对time进行格式化
	CreatedAt Timestamp `gorm:"column:created_at" json:"created_at"`
	UpdatedAt Timestamp `gorm:"column:created_at" json:"updated_at"`
}

// TableName 指定模型的表名称
func (User) TableName() string {
	return TableUser
}
