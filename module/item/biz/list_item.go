package biz

import (
	"context"
	"social_todo/common"
	"social_todo/module/item/model"
)

func (biz *itemBusiness) ListItems(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	data, err := biz.storage.ListItems(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
