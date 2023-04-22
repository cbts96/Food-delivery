package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(context context.Context, 
	condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) { //map[string] la giá tri dang key value trong bảng database ma ta can search, co interface la gia tri co kieu gi cung dc.
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil { // nen dung first vi chi tim kiem 1 cái, perfomance hon
		return nil, err
	}
	return &data, nil
}
