package biz

import (
	"context"
	"social_todo/module/item/model"
	"social_todo/module/item/utils"
)

func (biz *itemBusiness) CreateItem(ctx context.Context, itemData *model.TodoItemCreation) error {
	if !itemData.Validate() {
		return utils.ErrTitleEmpty
	}
	return biz.storage.CreateItem(ctx, itemData)
}
