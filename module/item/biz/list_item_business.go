package biz

import (
	"context"
	"social_todo/common"
	"social_todo/module/item/model"
)

type ListItemStorage interface {
	ListItems(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error)
}

type listItemBusiness struct {
	storage   ListItemStorage
	requester common.Requester
}

func NewListItemBusiness(storage ListItemStorage, requester common.Requester) *listItemBusiness {
	return &listItemBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *listItemBusiness) ListItems(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	ctx = context.WithValue(ctx, common.CurrrentUser, biz.requester)

	data, err := biz.storage.ListItems(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
