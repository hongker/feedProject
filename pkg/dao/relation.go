package dao

import (
	"feedProject/pkg/entity"
	"feedProject/pkg/enum"
	"fmt"
	"github.com/jinzhu/gorm"
)

// RelationDao
type RelationDao struct {
	BaseDao
}

// Relation
func Relation(db *gorm.DB) *RelationDao {
	dao := &RelationDao{}
	dao.db = db
	return dao
}

// GetByUserIdAndTargetId 根据UserId和TargetId获取实体
func (dao *RelationDao) GetByUserIdAndTargetId(userId, targetId int) (*entity.Relation, error) {
	result := new(entity.Relation)
	query := dao.db.Table(entity.TableRelation).Where("user_id = ? and target_id = ?", userId, targetId)

	if err := query.First(&result).Error;err != nil {
		return nil, err
	}

	return result, nil
}

// Update 更新
func (dao *RelationDao) Update(id int, columns entity.Columns) error {
	return dao.db.Table(entity.TableRelation).Where("id = ?", id).Updates(columns).Error
}

// GetFollowUserIds 获取粉丝ID
func (dao *RelationDao) GetFollowUserIds(userId int) ([]int, error) {
	var ids []int
	query := dao.db.Table(entity.TableRelation).
		Where("target_id = ? and status = ?", userId, enum.RelationStatusEnable)

	if err := query.Pluck("user_id", &ids).Error; err != nil {
		return nil, err
	}

	return ids, nil
}

// GetActiveFollowUserIds 获取活跃的粉丝ID
func (dao *RelationDao) GetActiveFollowUserIds(userId int) ([]int , error){
	var ids []int
	query := dao.db.Table(fmt.Sprintf("%s as r", entity.TableRelation)).
		Joins(fmt.Sprintf("left join %s as u on r.user_id=u.id", entity.TableUser)).
		Where("r.target_id = ? and r.status = ? and u.status = ?", userId, enum.RelationStatusEnable, enum.UserStatusActive)

	if err := query.Pluck("r.user_id", &ids).Error; err != nil {
		return nil, err
	}

	return ids, nil
}