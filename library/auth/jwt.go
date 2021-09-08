package auth

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/golang-jwt/jwt"
	_errors "github.com/pkg/errors"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = _errors.New("Token is expired")
	TokenNotValidYet = _errors.New("Token not active yet")
	TokenMalformed   = _errors.New("That's not even a token")
	TokenInvalid     = _errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{[]byte(global.Config.Jwt.SigningKey)}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims *request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims *request.CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
