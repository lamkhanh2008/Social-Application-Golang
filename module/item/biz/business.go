package biz

import (
	"context"
	"social_todo/common"
	"social_todo/module/item/model"
)

type itemBusiness struct {
	storage ItemStorageInterface
}

type ItemStorageInterface interface {
	CreateItem(ctx context.Context, itemData *model.TodoItemCreation) error
	GetByID(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteById(ctx context.Context, cond map[string]interface{}) error
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
	ListItems(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error)
}

func NewItemBusiness(storage ItemStorageInterface) *itemBusiness {
	return &itemBusiness{
		storage: storage,
	}
}
