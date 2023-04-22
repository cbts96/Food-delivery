package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db.Where("status in (1)")
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id=?", f.OwnerId)
		}
	}
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
