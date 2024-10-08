package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/component/tokenprovider"

	"log"
	"time"
)

type jwtProvider struct {
	prefix string
	secret string
}

func NewTokenJwtProvider(prefix string, secret string) *jwtProvider {
	return &jwtProvider{prefix: prefix, secret: secret}
}

type myClaims struct {
	Payload appCommon.TokenPayload `json:"payload"`
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

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		appCommon.TokenPayload{
			Uid:   data.UserId(),
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
		log.Println(err.Error())
		return nil, err
	}

	// return the token
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil

}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
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
