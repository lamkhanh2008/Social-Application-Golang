package biz

import (
	"context"
	"social_todo/module/item/model"
	"social_todo/module/item/utils"
)

func (biz *itemBusiness) UpdateItem(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	if !dataUpdate.Validate() {
		return utils.ErrTitleEmpty
	}

	data, err := biz.GetItemById(ctx, id)
	if err != nil {
		return err
	}

	if data.Status == "Deleted" {
		return err
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
