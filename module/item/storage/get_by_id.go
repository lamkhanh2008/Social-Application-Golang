package storage

import (
	"context"
	"social_todo/common"
	"social_todo/module/item/model"

	"gorm.io/gorm"
)

func (sqlStore *itemStorage) GetByID(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sqlStore.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(model.EntityName, err)
		}

		return nil, common.ErrDB(err)

	}

	return &data, nil
}
