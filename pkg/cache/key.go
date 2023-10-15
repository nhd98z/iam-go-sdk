package cache

import "strings"

const (
	KeyJWKPrefix = "jwk"

	KeyDelimiter = ":"
)

func Key(parts ...string) string {
	return strings.Join(parts, KeyDelimiter)
}
