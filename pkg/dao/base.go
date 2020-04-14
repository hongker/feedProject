package dao

import "github.com/jinzhu/gorm"

// BaseDao 基础Dao
type BaseDao struct {
	db *gorm.DB
}

// Create 创建
func (dao *BaseDao) Create(entity interface{}) error  {
	return dao.db.Omit("created_at", "updated_at").
		Create(entity).Error
}
