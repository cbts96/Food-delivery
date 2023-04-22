package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)
func DeleteRestaurant(appCtx appctx.AppContext) func (c *gin.Context){
	return func(c *gin.Context){
		db:= appCtx.GetMainDBCollection();
		id,err:=strconv.Atoi(c.Param("id"))
		if err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		store :=restaurantstorage.NewSqlStore(db)
		business:= restaurantbusiness.NwDeleteRestaurantBusiness(store)
		if err :=business.DeleteRestaurant(c.Request.Context(),id);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		// db.Table(Restaurant{}.TableName()).Where("id=?",id).Delete(nil)
		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}