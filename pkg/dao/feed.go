package dao

import (
	"feedProject/pkg/entity"
	"feedProject/pkg/enum"
	"feedProject/pkg/request"
	"fmt"
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

func (dao *FeedDao) Get(id int) (*entity.Feed, error) {
	feed := entity.Feed{}
	if err := dao.db.Table(entity.TableFeed).Where("id = ?", id).First(&feed).Error; err != nil {
		return nil, err
	}

	return &feed, nil
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


// GetFollowUserFeedItems 分页获取关注对象的feed
func (dao *FeedDao) GetFollowUserFeedItems(userId int, lastId, limit int) ([]entity.Feed, error) {
	query := dao.db.Table(fmt.Sprintf("%s as f", entity.TableFeed)).
		Joins(fmt.Sprintf("left join %s as r on f.creator_id=r.target_id", entity.TableRelation)).
		Where("r.user_id =? and r.status = ?", userId, enum.RelationStatusEnable).
		Where("f.status = ? and f.id > ?", enum.FeedStatusEnable, lastId).
		Order("f.id asc"). // 这里很重要，因为再表数据特别大的适合，需要用 f.id > lastId来获取分页数据，所以需要正序排列
		Limit(limit)

	var feeds []entity.Feed

	if err := query.Select("f.*").Scan(&feeds).Error; err != nil {
		return nil, err
	}

	return feeds, nil
}
