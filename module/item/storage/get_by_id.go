package storage

import (
	"context"
	"social_todo/module/item/model"
	"social_todo/module/item/utils"

	"gorm.io/gorm"
)

func (sqlStore *itemStorage) GetByID(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := sqlStore.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrNotFound
		}
	}

	return &data, nil
}
