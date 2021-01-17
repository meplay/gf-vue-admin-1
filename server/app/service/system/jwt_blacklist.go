package service

import (
	"database/sql"
	"errors"
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
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
	if errors.Is(b.db.Where("jwt = ?", jwt).Struct(&model.JwtBlacklist{}), sql.ErrNoRows) {
		return false
	}
	return true
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取用户在Redis的token
func (b *blacklist) GetRedisJWT(userUUID string) (string, error) {
	conn := g.Redis().Conn()
	defer func() {
		conn.Close()
	}()
	r, err := conn.Do("GET", userUUID)
	return gconv.String(r), err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 保存jwt到Redis
func (b *blacklist) SetRedisJWT(userUUID string, jwt string) error {
	_, err := g.Redis().Do("SETEX", userUUID, g.Cfg("jwt").GetUint("jwt.ExpiresAt")*3600000000000, jwt)
	return err
}
