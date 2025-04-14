package storage

import (
	"context"
	"social_todo/common"
	"social_todo/module/item/model"
)

func (sqlStorage *itemStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := sqlStorage.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
