package storage

import (
	"context"
	"social_todo/common"
	"social_todo/module/user/model"
)

func (store *sqlStore) CreateUser(ctx context.Context, user *model.UserCreate) error {
	db := store.db.Begin()

	if err := db.Table(user.TableName()).Create(user).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
