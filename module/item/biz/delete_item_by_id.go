package biz

import "context"

func (biz *itemBusiness) DeleteItemById(ctx context.Context, id int) error {
	if _, err := biz.GetItemById(ctx, id); err != nil {

	}

	err := biz.storage.DeleteById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	return nil
}
