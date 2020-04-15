package dao

import (
	"feedProject/pkg/entity"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	BaseDao
}

func User(db *gorm.DB) *UserDao {
	dao := &UserDao{}
	dao.db = db
	return dao
}

func (dao *UserDao) Get(id int) (*entity.User, error) {
	user := entity.User{}
	if err := dao.db.Table(entity.TableUser).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}