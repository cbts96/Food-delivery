package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) Delete(context context.Context, id int ) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
	Where("id=?",id).Updates(map[string]interface{}{"status":0}).Error; err!=nil{ //soft delete: dung updates thay vi delete vi nếu delete là delete hết, nếu có 1 danh muc có nhiều p.tu con, neu xoa het se lam loi các phần tử liên quan , // status 0 là da delete, còn 1 là create
		return err
	}
	return nil
}