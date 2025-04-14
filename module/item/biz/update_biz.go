package biz

import (
	"context"
	"errors"
	"social_todo/common"
	"social_todo/module/item/model"
	"social_todo/module/item/utils"
)

type UpdateItemStorage interface {
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
	GetByID(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type updateItemBusiness struct {
	storage   UpdateItemStorage
	requester common.Requester
}

func NewUpdateItemBusiness(storage UpdateItemStorage, requester common.Requester) *updateItemBusiness {
	return &updateItemBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *updateItemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.storage.GetByID(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz *updateItemBusiness) UpdateItem(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {

	if !dataUpdate.Validate() {
		return common.ErrValidate(utils.ErrTitleEmpty)
	}

	data, err := biz.GetItemById(ctx, id)
	if err != nil {
		return err
	}

	if data.Status == "Deleted" {
		return err
	}

	isOwner := data.UserId != biz.requester.GetUserId()
	if !isOwner && !common.IsAdmin(biz.requester) {
		return common.ErrNoPermission(errors.New("No permission"))
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
