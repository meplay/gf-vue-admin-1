package request

import "github.com/dgrijalva/jwt-go"

// CustomClaims Custom claims structure
type CustomClaims struct {
	ID          uint
	BufferTime  int64
	UUID        string
	Username    string
	NickName    string
	AuthorityId string

	jwt.StandardClaims
}
