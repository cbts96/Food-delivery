package ginrestaurant

import (
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func CreateRestaurant(db *gorm.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error" :err.Error(),
			})
			return
		}
		store:= restaurantstorage.NewSqlStore(db)
		business :=  restaurantbusiness.NwCreateResBusiness(store)
		// business.CreateRestaurant(c.Request.Context())
		
		if err:= business.CreateRestaurant(c.Request.Context(),&data); err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"data": data,
		})
	}
}