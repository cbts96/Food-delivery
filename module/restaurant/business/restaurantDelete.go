package restaurantbusiness

import (
	"context"
	
	// restaurantmodel "food-delivery/module/restaurant/model"
)
type DeleteRestaurantStore interface {
	Delete(context context.Context, id int) error
}
type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness{
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(context context.Context, id int) error{


	
	if  err := business.store.Delete(context,id); err!=nil{
		return err
	}
	return nil
}