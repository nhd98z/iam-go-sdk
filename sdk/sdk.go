package sdk

import (
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/entity"
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/token/jwt"
)

type SDK struct{}

func New() *SDK {
	return &SDK{}
}

func (s *SDK) ParseBearerJWT(bearerTokenJWTString string) (*entity.AccessToken, error) {
	return jwt.ParseBearer(bearerTokenJWTString)
}

func (s *SDK) ParseJWT(tokenJWTString string) (*entity.AccessToken, error) {
	return jwt.Parse(tokenJWTString)
}
