package jwt

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"

	"github.com/golang-jwt/jwt/v5"

	"iam-go-sdk/pkg/cache"
	"iam-go-sdk/pkg/oauth/client"
	"iam-go-sdk/pkg/oauth/entity"
)

func ParseBearer(bearerAuthorization string) (*jwt.Token, error) {
	var bearerSchema = "Bearer "
	if len(bearerAuthorization) <= len(bearerSchema) {
		return nil, fmt.Errorf("encountered error when parsing bearer jwt: invalid bearer authorization")
	}
	return ParseBearer(bearerAuthorization[len(bearerSchema):])
}

func Parse(tokenJwt string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenJwt,
		func(token *jwt.Token) (interface{}, error) {
			if _, isValid := token.Method.(*jwt.SigningMethodRSA); !isValid {
				return nil, fmt.Errorf("encountered error when parsing jwt: invalid signing method: %s", token.Header["alg"])
			}

			kid := token.Header["kid"].(string)
			if kid == "" {
				return nil, errors.New("encountered error when parsing jwt: invalid token kid")
			}

			jwk, err := getJWKByKid(kid)
			if err != nil {
				return nil, err
			}

			decodedE, err := base64.RawURLEncoding.DecodeString(jwk.E)
			if err != nil {
				return nil, fmt.Errorf("encountered error when parsing jwt while decoding E: %v", err)
			}
			bigE := new(big.Int).SetBytes(decodedE)

			decodedN, err := base64.RawURLEncoding.DecodeString(jwk.N)
			if err != nil {
				return nil, fmt.Errorf("encountered error when parsing jwt while decoding N: %v", err)
			}

			bigN := new(big.Int).SetBytes(decodedN)
			return &rsa.PublicKey{
				N: bigN,
				E: int(bigE.Int64()),
			}, nil
		})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("encountered error when parsing jwt: token is invalid")
	}

	return token, nil
}

func getJWKByKid(kid string) (*entity.JWK, error) {
	// check cache
	cacheKey := cache.Key(cache.KeyJWKPrefix, kid)
	cached, err := cache.Get[*entity.JWK](cacheKey)
	if err == nil && cached != nil {
		return cached, nil
	}

	// get from OAuth API
	jwks, err := client.GetJWKs()
	if err != nil {
		return nil, err
	}
	jwk := jwks[kid]
	if jwk == nil {
		return nil, fmt.Errorf("encountered error when fetching jwk: not found kid %s", kid)
	}

	// set cache
	cache.Set(cacheKey, jwk)

	return jwk, nil
}
