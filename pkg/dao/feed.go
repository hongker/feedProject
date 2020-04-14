package dao

import (
	"feedProject/pkg/entity"
	"feedProject/pkg/enum"
	"feedProject/pkg/request"
	"github.com/jinzhu/gorm"
)

type FeedDao struct {
	BaseDao
}

// Feed
func Feed(db *gorm.DB) *FeedDao {
	dao := &FeedDao{}
	dao.db = db
	return dao
}

func (dao *FeedDao) Query() {

}

// Update 更新
func (dao *FeedDao) Update(id int, columns entity.Columns) error {
	return dao.db.Table(entity.TableFeed).Where("id = ?", id).Updates(columns).Error
}

// SearchFeed
func (dao *FeedDao) SearchFeed(req request.FeedSearchRequest) ([]entity.Feed, error) {
	var result []entity.Feed
	query := dao.db.Table(entity.TableFeed).
		Where("status = ?", enum.FeedStatusEnable).
		Where("creator_id in (?)", req.CreatorIds).
		Where("id > ?", req.LastPullId).
		Order("id asc")

	if err := query.Limit(req.Limit).Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}
