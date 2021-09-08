package system

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/response"
	"github.com/flipped-aurora/gf-vue-admin/library/auth"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/go-redis/redis/v8"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

var User = new(user)

type user struct{}

// Register
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *user) Register(info *request.UserRegister) (data *system.User, err error) {
	var entity system.User
	if !errors.Is(global.Db.Where("username = ?", info.Username).First(&entity).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return nil, _errors.Wrap(err, "用户名已注册")
	}
	create := info.Create()
	if err = entity.EncryptedPassword(); err != nil {
		return nil, _errors.Wrap(err, "密码加密失败!")
	}
	if err = global.Db.Create(&create).Error; err != nil {
		return nil, _errors.Wrap(err, "用户注册失败!")
	}
	return &create, nil
}

// Login 用户登录接口
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *user) Login(info *request.UserLogin) (data *response.UserLogin, err error) {
	var entity system.User
	if errors.Is(global.Db.Where("username = ?", info.Username).Preload("Authority").First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil, _errors.New("用户不存在!")
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return nil, _errors.New("密码错误!")
	}
	return s.tokenNext(&entity)
}

// Find 查找用户详细信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *user) Find(info *request.UserFind) (data *system.User, err error) {
	var entity system.User
	if err = global.Db.Scopes(info.Search()).Preload("Authorities").Preload("Authority").First(&entity).Error; err != nil {
		return nil, _errors.Wrap(err, "用户查询失败!")
	}
	return &entity, nil
}

// Update 更新用户信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *user) Update(info *request.UserUpdate) (data *system.User, err error) {
	update := info.Update()
	err = global.Db.Where("id = ?", info.ID).Updates(&update).Error
	return &update, err
}

// ChangePassword 修改密码
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *user) ChangePassword(info *request.UserChangePassword) error {
	var entity system.User
	err := global.Db.Where("username = ?", info.Username).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return _errors.Wrap(err, "用户不存在! ")
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return _errors.Wrap(err, "密码错误!")
	}
	entity.Password = info.NewPassword
	if err = entity.EncryptedPassword(); err != nil {
		return _errors.Wrap(err, "密码加密失败!")
	}
	return global.Db.Where("username = ?", info.Username).Update("password", entity.Password).Error
}

// SetAuthority 设置用户的活跃角色
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *user) SetAuthority(info *request.UserSetAuthority) error {
	err := global.Db.Where("user_id = ? AND authority_id = ?", info.Uuid, info.AuthorityId).First(&system.UseAuthority{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return _errors.New("该用户无此角色!")
	}
	if err = global.Db.Model(&system.User{}).Where("uuid = ?", info.Uuid).Update("authority_id", info.AuthorityId).Error; err != nil {
		return _errors.Wrap(err, "更新用户角色失败!")
	}
	return nil
}

// SetUserAuthorities 设置用户可切换的角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *user) SetUserAuthorities(info *request.UserSetAuthorities) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&[]system.UseAuthority{}, "user_id = ?", info.ID).Error; err != nil {
			return _errors.Wrap(err, "用户可切换的旧角色删除失败!")
		}
		length := len(info.AuthorityIds)
		entities := make([]system.UseAuthority, 0, length)
		for i := 0; i < length; i++ {
			entities = append(entities, system.UseAuthority{UserId: info.ID, AuthorityId: info.AuthorityIds[i]})
		}
		if err := tx.Create(&entities).Error; err != nil {
			return _errors.Wrap(err, "设置用户多角色失败!")
		}
		return nil
	})
}

// Delete 删除用户
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *user) Delete(info *common.GetByID) error {
	return global.Db.Delete(&system.User{}, info.ID).Error
}

// GetList 获取用户列表
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *user) GetList(info *common.PageInfo) (list []system.User, total int64, err error) {
	entities := make([]system.User, 0, info.PageSize)
	db := global.Db.Model(&system.User{})
	err = db.Count(&total).Error
	err = db.Scopes(common.Paginate(info)).Preload("Authority").Find(&entities).Error
	return entities, total, err
}

// tokenNext 生成token
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *user) tokenNext(user *system.User) (*response.UserLogin, error) {
	_jwt := auth.NewJWT()
	claims := request.CustomClaims{
		Uuid:        user.Uuid,
		ID:          user.ID,
		Nickname:    user.Nickname,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.Config.Jwt.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Config.Jwt.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                          // 签名的发行者
		},
	}
	token, err := _jwt.CreateToken(&claims)
	if err != nil {
		return nil, _errors.Wrap(err, "获取token失败!")
	}
	if !global.Config.System.UseMultipoint {
		entity := response.UserLogin{User: user, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
		return &entity, nil
	}

	if jwtStr, _err := JwtBlacklist.GetRedisJWT(user.Username); _err == redis.Nil {
		if err = JwtBlacklist.SetRedisJWT(token, user.Username); err != nil {
			return nil, _errors.Wrap(err, "设置登录状态失败!")
		}
		entity := response.UserLogin{User: user, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
		return &entity, nil
	} else if _err != nil {
		return nil, _errors.Wrap(_err, "设置登录状态失败!")

	} else {
		if !JwtBlacklist.IsBlacklist(jwtStr) {
			return nil, _errors.Wrap(_err, "jwt作废失败!")
		}
		if err = JwtBlacklist.SetRedisJWT(token, user.Username); err != nil {
			return nil, _errors.Wrap(err, "设置登录状态失败!")
		}
		entity := response.UserLogin{User: user, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
		return &entity, nil
	}
}
