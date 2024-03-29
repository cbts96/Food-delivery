package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc{
	return func(c *gin.Context){
		db := appCtx.GetMainDBCollection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		pagingData.Fulfill()
		var filter restaurantmodel.Filter
		if err :=c.ShouldBind(&filter);err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}
		store := restaurantstorage.NewSqlStore(db)
		business := restaurantbusiness.NewListRestaurantBusiness(store)
		
		result,err := business.ListRestaurant(c.Request.Context(), &filter ,&pagingData)
		
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,common.NewSuccessResponse(&result,pagingData,filter))
	}
}