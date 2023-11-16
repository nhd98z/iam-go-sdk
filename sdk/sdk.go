package sdk

import (
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/entity"
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/token/jwt"
	permissionclient "github.com/anhvietnguyennva/iam-go-sdk/permission/client"
	"github.com/anhvietnguyennva/iam-go-sdk/permission/dto"
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

func (s *SDK) CheckPermission(namespace string, object string, relation string, subjectID string, maxDepth uint8) (bool, error) {
	return permissionclient.CheckPermission(namespace, object, relation, subjectID, maxDepth)
}

func (s *SDK) CreatePermission(request *dto.CreatePermissionRequest, bearerAccessToken string) (string, error) {
	return permissionclient.CreatePermission(request, bearerAccessToken)
}
