package entity

import "github.com/golang-jwt/jwt/v5"

type AccessToken struct {
	jwt.Token

	ClientID       string
	Subject        string
	ExpirationTime int64
	IssueAt        int64
	Scopes         []string
}
