package storage

import (
	"context"
	"social_todo/module/item/model"
)

func (sqlStore *itemStorage) DeleteById(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "Deleted"
	if err := sqlStore.db.Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus,
		}).Error; err != nil {
		return err
	}

	return nil
}
