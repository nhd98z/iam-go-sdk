package sdk

import (
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/entity"
	"github.com/anhvietnguyennva/iam-go-sdk/permission/dto"
)

type ISDK interface {
	ParseBearerJWT(bearerTokenJWTString string) (*entity.AccessToken, error)
	ParseJWT(tokenJWTString string) (*entity.AccessToken, error)
	CheckPermission(namespace string, object string, relation string, subjectID string, maxDepth uint8) (bool, error)
	CreatePermission(request *dto.CreatePermissionRequest, bearerAccessToken string) (string, error)
}
