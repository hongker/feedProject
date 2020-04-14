package dao

import "github.com/jinzhu/gorm"

type UserDao struct {
	BaseDao
}

func User(db *gorm.DB) *UserDao {
	dao := &UserDao{}
	dao.db = db
	return dao
}

