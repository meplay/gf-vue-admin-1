package service

import (
	"context"
	"errors"
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"gorm.io/gorm"
	"time"
)

var JwtBlacklist = new(blacklist)

type blacklist struct {
	_blacklist model.JwtBlacklist
}

// JwtToBlacklist 拉黑jwt
// Author [Aizen1172](https://github.com/Aizen1172)
func (b *blacklist) JwtToBlacklist(jwt string) error {
	entity := model.JwtBlacklist{Jwt: jwt}
	return global.Db.Create(&entity).Error
}

// IsBlacklist 判断JWT是否在jwt黑名单
// Author [Aizen1172](https://github.com/Aizen1172)
func (b *blacklist) IsBlacklist(jwt string) bool {
	return !errors.Is(global.Db.Where("jwt = ?",jwt).First(&model.JwtBlacklist{}).Error,gorm.ErrNotImplemented)
}

// GetRedisJWT 获取用户在Redis的token
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *blacklist) GetRedisJWT(uuid string) (string, error) {
	return global.Redis.Get(context.Background(), uuid).Result()
}

// SetRedisJWT 保存jwt到Redis
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *blacklist) SetRedisJWT(uuid string, jwt string) error {
	timer := time.Duration(global.Config.Jwt.ExpiresAt) * time.Second
	return global.Redis.Set(context.Background(), uuid, jwt, timer).Err()
}

// ValidatorRedisToken 鉴权jwt
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *blacklist) ValidatorRedisToken(userUUID string, oldToken string) bool {
	if jwt, err := b.GetRedisJWT(userUUID); err != nil {
		return false
	} else {
		if jwt != oldToken {
			return false
		}
		return true
	}
}
