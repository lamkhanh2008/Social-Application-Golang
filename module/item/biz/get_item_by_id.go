package biz

import (
	"context"
	"social_todo/module/item/model"
)

func (biz *itemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.storage.GetByID(ctx, map[string]interface{}{"id": id})
	if err != nil {

	}

	return data, nil
}
