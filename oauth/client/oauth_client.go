package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	constant2 "github.com/anhvietnguyennva/iam-go-sdk/constant"
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/dto"
	"github.com/anhvietnguyennva/iam-go-sdk/oauth/entity"
	"github.com/anhvietnguyennva/iam-go-sdk/util/env"
)

func GetJWKs() (map[string]*entity.JWK, error) {
	url := env.StringFromEnv(constant2.EnvKeyOAuthGetJWKsURL, constant2.OAuthGetJWKsDefaultURL)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bodyJson, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("encountered error when calling %s: %v", url, err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("encountered error when calling %s, status code: %d, response: %s", url, res.StatusCode, bodyJson)
	}
	var responseBody dto.GetJWKsResponse
	err = json.Unmarshal(bodyJson, &responseBody)
	if err != nil {
		return nil, fmt.Errorf("encountered error when attempting unmarshal response from %s: %v", url, err)
	}

	jwks := responseBody.ToEntities()
	ret := make(map[string]*entity.JWK, len(jwks))
	for _, jwk := range jwks {
		ret[jwk.Kid] = jwk
	}
	return ret, nil
}
