package client

import (
	"encoding/json"
	"fmt"
	"iam-go-sdk/pkg/oauth/dto"
	"io"
	"net/http"

	"iam-go-sdk/pkg/constant"
	"iam-go-sdk/pkg/oauth/entity"
	"iam-go-sdk/pkg/util/env"
)

func GetJWKs() (map[string]*entity.JWK, error) {
	url := env.StringFromEnv(constant.EnvKeyOAuthGetJWKsURL, constant.OAuthGetJWKsDefaultURL)
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
