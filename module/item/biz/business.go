package biz

import (
	"context"
	"social_todo/module/item/model"
)

type itemBusiness struct {
	storage ItemStorageInterface
	// requester common.Requester
}

type ItemStorageInterface interface {
	CreateItem(ctx context.Context, itemData *model.TodoItemCreation) error
	GetByID(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteById(ctx context.Context, cond map[string]interface{}) error
}

func NewItemBusiness(storage ItemStorageInterface) *itemBusiness {
	return &itemBusiness{
		storage: storage,
		// requester: requester,
	}
}
