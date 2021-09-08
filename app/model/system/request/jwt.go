package request

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          uint
	Uuid        string
	Nickname    string
	Username    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
