package storage

import "gorm.io/gorm"

type itemStorage struct {
	db *gorm.DB
}

func NewItemStorage(db *gorm.DB) *itemStorage {
	return &itemStorage{db: db}
}
