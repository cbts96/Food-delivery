package appctx

import "gorm.io/gorm"
type AppContext interface{
	GetMainDBCollection() *gorm.DB
}
type appCtx struct{
	db *gorm.DB
}
func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}
func (ctx *appCtx) GetMainDBCollection() *gorm.DB {return ctx.db}