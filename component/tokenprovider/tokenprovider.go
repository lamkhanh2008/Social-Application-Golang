package tokenprovider

import (
	"errors"
	"social_todo/common"
)

type Provider interface {
	Generate(data TokenPayLoad, expiry int) (Token, error)
	Validate(token string) (TokenPayLoad, error)
	SecretKey() string
}

type TokenPayLoad interface {
	UserId() int
	Role() string
}

type Token interface {
	GetToken() string
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)
	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"errEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(errors.New("Invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken")
)
