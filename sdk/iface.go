package sdk

import (
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/entity"
)

type ISDK interface {
	ParseBearerJWT(bearerTokenJWTString string) (*entity.AccessToken, error)
	ParseJWT(tokenJWTString string) (*entity.AccessToken, error)
}
