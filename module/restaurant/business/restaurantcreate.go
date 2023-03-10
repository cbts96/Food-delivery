package restaurantbusiness

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func NwCreateResBusiness(store CreateRestaurantStore) *createRestaurantBusiness { // phai co interface CreateRestaurantStore vì phai tach rieng ra, nếu ko kho unit test va pha vo cleanArchi....
	return &createRestaurantBusiness{store: store}
}
func (business *createRestaurantBusiness) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	//trong Go ko co try catch nen phai if else 2 lan
	if data.Name == "" {
		return errors.New(("Field name can't be empty!"))
	}

	if error := business.store.CreateRestaurant(context, data); error != nil {
		return error
	}
	return nil
}
