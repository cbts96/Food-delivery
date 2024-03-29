package restaurantbusiness

import (
	"context"
	"errors"

	restaurantmodel "food-delivery/module/restaurant/model"
)
type DeleteRestaurantStore interface {
	FindDataWithCondition(context context.Context, 
		condition map[string]interface{}, moreKeys ...string,) (*restaurantmodel.Restaurant, error); Delete(context context.Context, id int ) error
}
type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NwDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness{
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(context context.Context, id int) error{
 oldData, err := business.store.FindDataWithCondition(context, map[string] interface{}{"id":id})
 if err !=nil{
	return err
 }

 if oldData.Status==0{
	return errors.New("data has been deleted")
 }
	
	if  err := business.store.Delete(context,id); err!=nil{
		return err
	}
	return nil
}