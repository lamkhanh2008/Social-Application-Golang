package biz

import (
	"context"
	"social_todo/common"
	"social_todo/component/tokenprovider"
	"social_todo/module/user/model"
)

type LogicStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type logicBusiness struct {
	storage       LogicStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLogicBusiness(storeUser LogicStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *logicBusiness {
	return &logicBusiness{
		storage:       storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *logicBusiness) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	user, err := biz.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	newHashed := biz.hasher.Hash(data.Password + user.Salt)
	if newHashed != user.Password {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	payload := &tokenprovider.TokenPayLoadImpl{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, err
}
