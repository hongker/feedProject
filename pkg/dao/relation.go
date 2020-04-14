package dao

import (
	"feedProject/pkg/entity"
	"feedProject/pkg/enum"
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

// GetTargetIdByUserIdAndOffset 获取关注对象
func (dao *RelationDao) GetTargetIdByUserIdAndOffset(userId int, offset, limit int) ([]int, error) {
	var targetIds []int
	query := dao.db.Table(entity.TableRelation).
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Where("user_id = ? and status = ?", userId, enum.RelationStatusEnable)

	if err := query.Pluck("target_id", &targetIds).Error; err != nil {
		return nil, err
	}

	return targetIds, nil
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
