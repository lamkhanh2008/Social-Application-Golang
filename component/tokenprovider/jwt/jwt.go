package jwt

import (
	"fmt"
	"social_todo/component/tokenprovider"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	prefix string
	secret string
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayLoad, expiry int) (tokenprovider.Token, error) {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		tokenprovider.TokenPayLoadImpl{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &token{
		Token:   myToken,
		Created: now,
		Expiry:  expiry,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayLoad, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return claims.Payload, nil
}

func NewTokenJWTProvider(prefix string, secretKey string) *jwtProvider {
	return &jwtProvider{prefix: prefix, secret: secretKey}
}

type myClaims struct {
	Payload tokenprovider.TokenPayLoadImpl `json:"payload"`
	jwt.StandardClaims
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}
