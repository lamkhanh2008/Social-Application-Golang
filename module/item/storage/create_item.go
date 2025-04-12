package storage

import (
	"context"
	"social_todo/module/item/model"
)

func (store *itemStorage) CreateItem(ctx context.Context, itemData *model.TodoItemCreation) error {
	if err := store.db.Create(&itemData).Error; err != nil {
		return err
	}

	return nil
}
