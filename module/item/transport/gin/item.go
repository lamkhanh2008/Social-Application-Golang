package gin

import (
	"context"
	"social_todo/module/item/model"
)

type ItemBusiness interface {
	CreateItem(ctx context.Context, itemData *model.TodoItemCreation) error
	DeleteItemById(ctx context.Context, id int) error
	GetItemById(ctx context.Context, id int) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error
}
type service struct {
	biz ItemBusiness
}

func NewItemService(itemBusiness ItemBusiness) *service {
	return &service{
		biz: itemBusiness,
	}
}
