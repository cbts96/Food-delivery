package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error { //context context.context là liên quan tới IO, giúp ta cancel hay tiếp tuc groutine
	if error := s.db.Create(&data).Error; error!=nil{

		return error
	}
	return nil
}
