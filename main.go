// // package main

// // import (
// // 	"log"
// // 	"os"

// // 	"gorm.io/driver/mysql"
// // 	"gorm.io/gorm"
// // )

// // type Restaurant struct {
// // 	Id   int    `json: "id" gorm:"column:id";`
// // 	Name string `json: "id" gorm:"column:name";`
// // 	Addr string `json: "addr" gorm:"column:addr";`
// // }

// // func (Restaurant) TableName() string { return "restaurant" }

// // type RestaurantUpdate struct {
// // 	Name *string `json:"id" gorm:"column:name";`
// // 	Addr *string `json:"addr" gorm:"column:addr";`
// // }

// // func (RestaurantUpdate) TableName() string {
// // 	return Restaurant{}.TableName()
// // }
// // func main() {
// // 	os.Setenv("MYSQL_CNT_STRING", "root:lexuanthang123@tcp(127.0.0.1:3306)/Food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
// // 	// os.Getenv("MYSQL_CNT_STRING")
// // 	dsn := os.Getenv("MYSQL_CNT_STRING")
// // 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// // 	if err != nil {
// // 		log.Fatalln(err)
// // 	}
// // 	log.Println(db, err)
// // 	newResta := Restaurant{Name: "thang", Addr: "tien giang"}

// // 	if err := db.Create(&newResta).Error; err != nil {

// // 		log.Println(err)
// // 	}
// // 	var myRestaurant Restaurant
// // 	if err := db.Where("?", 3).First(&myRestaurant).Error; err != nil {
// // 		log.Println("err=>", err)
// // 	}
// // 	newName := "lethang"
// // 	updateName := RestaurantUpdate{Name: &newName}

// // 	// log.Println(myRestaurant)
// // 	myRestaurant.Name = "Thang Restaurant"
// // 	if err := db.Where("?", 1).Updates(&updateName).Error; err != nil {
// // 		log.Println("err=>", err)
// // 	}
// // 	// log.Println(myRestaurant)
// // 	if err := db.Table(Restaurant{}.TableName()).Where("?", 1).Delete(nil).Error; err != nil {
// // 		log.Println(err)
// // 	}
// // 	log.Println(myRestaurant)
// // }

// // // // main()

// // // func main ()  {

// // // 	// os.Getenv("MYSQL_CNT_STRING")
// // // 	dsn := "root:lexuanthang123@tcp(127.0.0.1:3306)/Food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
// // //   	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// // // 	// if err !=nil{
// // // 	// 	log.Fatalln(err)
// // // 	// }
// // // 	log.Println("123",db,err)
// // // }
// // // // main()
package main

import (
	"food-delivery/component/appctx"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"log"
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "strconv"
)

type Restaurant struct {
	Id   int    `json: "id" gorm:"column:id";`
	Name string `json: "id" gorm:"column:name";`
	Addr string `json: "addr" gorm:"column:addr";`
}

func (Restaurant) TableName() string { return "restaurant" }

type RestaurantUpdate struct {
	Name *string `json:"id" gorm:"column:name";`
	Addr *string `json:"addr" gorm:"column:addr";`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
func main() {
	os.Setenv("MYSQL_CNT_STRING", "root:lexuanthang123@tcp(127.0.0.1:3306)/Food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
	// os.Getenv("MYSQL_CNT_STRING")

	dsn := os.Getenv("MYSQL_CNT_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	appContext := appctx.NewAppContext(db)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)
	r := gin.Default()
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))
	// v1.GET("/restaurants/:id",ginrestaurant.)
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// package main

// import (
// 	"food-delivery/module/restaurant/transport/ginrestaurant"
// 	"log"
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Restaurant struct{
// 	Id int `json: "id" gorm:"column:id";`
// 	Name string `json: "name" gorm:"column:name";`
// 	Addr string `json: "addr" gorm:"column:addr";`
// }

// func (Restaurant) TableName() string{
// 	return "restaurant"
// }

// type RestaurantUpdate struct{
// 	Name *string `json:"name" gorm:"column:name;"`
// 	Addr *string `json:"addr" gorm:"column:addr;"`
// }

// func main(){
// 	os.Setenv("MYSQL_CNT_STRING", "root:lexuanthang123@tcp(127.0.0.1:3306)/Food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
// 	dsn:= os.Getenv("MYSQL_CNT_STRING")
// 	db, err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil{
// 		log.Fatalln(err)
// 	}
// 	log.Println(db,err)
// 	r:=gin.Default()
// 	v1:= r.Group(("/v1"))
// 	v1.POST("/restuarants", ginrestaurant.CreateRestaurant(db))
// }
