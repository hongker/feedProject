package service

import (
	"feedProject/pkg/app"
	"feedProject/pkg/dao"
	"feedProject/pkg/entity"
	"feedProject/pkg/enum"
	"fmt"
	"github.com/jinzhu/gorm"
)

// RelationService 关注逻辑
type RelationService struct {
}

func Relation() *RelationService {
	return &RelationService{}
}

// Create 创建关注
func (service *RelationService) Create(userId, targetId int) error {
	// TODO 缓存设计
	relationDao := dao.Relation(app.DB())

	// check has create relation
	relation, err := relationDao.GetByUserIdAndTargetId(userId, targetId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("GetByUserIdAndTargetId:%v", err)
	}

	if relation == nil {
		// create
		if err := relationDao.Create(&entity.Relation{
			UserId:   userId,
			TargetId: targetId,
			Status: enum.FeedStatusEnable,
		}); err != nil {
			return fmt.Errorf("create: %v", err)
		}
	} else {
		// update
		if err := relationDao.Update(relation.ID, entity.Columns{
			"status": enum.RelationStatusEnable,
		}); err != nil {
			return fmt.Errorf("update: %v", err)
		}
	}
	return nil
}

// Cancel 取消关注
func (service *RelationService) Cancel(userId, targetId int) error {
	relationDao := dao.Relation(app.DB())

	// check has create relation
	relation, err := relationDao.GetByUserIdAndTargetId(userId, targetId)
	if err != nil {
		return fmt.Errorf("GetByUserIdAndTargetId:%v", err)
	}

	// update
	if err := relationDao.Update(relation.ID, entity.Columns{
		"status": enum.RelationStatusDisable,
	}); err != nil {
		return fmt.Errorf("update: %v", err)
	}

	return nil
}
