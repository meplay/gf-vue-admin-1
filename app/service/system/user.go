package system

import (
	"errors"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

var User = new(user)

type user struct {}

func (s *user) Login(info *request.Login) (data *system.User, err error) {
	var entity system.User
	if errors.Is(global.Db.Where("username = ?", info.Username).Preload("Authority").First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil, _errors.New("用户不存在")
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return &entity, _errors.New("密码错误")
	}
	return &entity, nil
}