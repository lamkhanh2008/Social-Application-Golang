package biz

import (
	"context"
	"social_todo/module/item/model"
)

func (biz *itemBusiness) CreateItem(ctx context.Context, itemData *model.TodoItemCreation) error {
	if err := itemData.Validate(); err != nil {
		return err
	}
	return biz.storage.CreateItem(ctx, itemData)
}
