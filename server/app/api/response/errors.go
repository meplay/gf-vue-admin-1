package response

import (
	"errors"
)

var (
	TokenExpired     = errors.New("Couldn't Handle This Token! ")
	TokenInvalid     = errors.New("Token Is Expired ! ")
	TokenMalformed   = errors.New("That's Not Even A Token! ")
	TokenNotValidYet = errors.New("Token Not Active Yet! ")

	ErrorRegister           = errors.New("Register Failed! ")
	ErrorUserNoExist        = errors.New("User Is No Exist! ")
	ErrorWrongPassword      = errors.New("Wrong Password! ")
	ErrorEncryptedPassword  = errors.New("Password Encryption Failed! ")
	ErrorUsernameRegistered = errors.New("Username Is Registered! ")

	ErrorSameAPI  = errors.New("存在相同api")
	ErrorSamePath = errors.New("存在相同api路径")

	ErrorSameType = errors.New("存在相同的type，不允许创建")
	ErrorSameAuthority = errors.New("存在相同角色id")
)
