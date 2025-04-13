package biz

import (
	"context"
	"social_todo/common"
	"social_todo/module/item/model"
)

func (biz *itemBusiness) DeleteItemById(ctx context.Context, id int) error {
	if _, err := biz.GetItemById(ctx, id); err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	err := biz.storage.DeleteById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	return nil
}
