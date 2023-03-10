package ginrestaurant

import (
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func DeleteRestaurant(db *gorm.DB) func (c *gin.Context){
	return func(c *gin.Context){
		id,err:=strconv.Atoi(c.Param("id"))
		if err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		store :=restaurantstorage.NewSqlStore(db)
		business:= restaurantbusiness.NewDeleteRestaurantBusiness(store)
		if err :=business.DeleteRestaurant(c.Request.Context(),id);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		// db.Table(Restaurant{}.TableName()).Where("id=?",id).Delete(nil)
		c.JSON(http.StatusOK,gin.H{
			"data":1,
		})
	}
}