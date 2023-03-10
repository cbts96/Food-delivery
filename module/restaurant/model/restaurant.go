package restaurantmodel

//can chia ra Restaurant / RestaurantCreate / RestaurantUpdate vì lien quan sercurity có 1 số vd: id, name,... là tt nhạy cam can che di, nên phai han chế 1 so trường, cũng như tiêt kiem bo nhớ
type Restaurant struct {
	Id   int    `json: "id" gorm:"column:id";`
	Name string `json: "id" gorm:"column:name";`
	Addr string `json: "addr" gorm:"column:addr";`
}

func (Restaurant) TableName() string { return "restaurant" }
type RestaurantCreate struct {
	Id   int    `json: "id" gorm:"column:id";`
	Name string `json: "id" gorm:"column:name";`
	Addr string `json: "addr" gorm:"column:addr";`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string `json:"id" gorm:"column:name";`
	Addr *string `json:"addr" gorm:"column:addr";`
}
