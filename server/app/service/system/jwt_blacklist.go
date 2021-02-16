package service

import (
	"context"
	"database/sql"
	"errors"
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"time"
)

var JwtBlacklist = &blacklist{db: g.DB().Table("jwt_blacklist").Safe()}

type blacklist struct {
	db *gdb.Model
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 拉黑jwt
func (b *blacklist) JwtToBlacklist(jwt string) error {
	var entity = model.JwtBlacklist{Jwt: jwt}
	_, err := b.db.Data(&entity).Insert()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 判断JWT是否在jwt黑名单
func (b *blacklist) IsBlacklist(jwt string) bool {
	return !errors.Is(b.db.Where("jwt = ?", jwt).Struct(&model.JwtBlacklist{}), sql.ErrNoRows)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取用户在Redis的token
func (b *blacklist) GetRedisJWT(uuid string) (string, error) {
	return global.Redis.Get(context.Background(), uuid).Result()
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 保存jwt到Redis
func (b *blacklist) SetRedisJWT(uuid string, jwt string) error {
	timer := time.Duration(global.Config.Jwt.ExpiresAt) * time.Second
	return global.Redis.Set(context.Background(), uuid, jwt, timer).Err()
}
