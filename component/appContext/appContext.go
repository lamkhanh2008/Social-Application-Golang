package appCtx

import "gorm.io/gorm"

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) AppContext {
	return &appContext{
		db: db,
	}
}

func (app *appContext) GetMaiDBConnection() *gorm.DB {
	return app.db
}
