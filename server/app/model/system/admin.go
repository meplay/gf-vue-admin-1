package model

import (
	"github.com/gogf/gf/os/gtime"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id          uint        `orm:"id,primary"   json:"ID"`          // 自增ID
	CreateAt    *gtime.Time `orm:"create_at"    json:"CreatedAt"`   // 创建时间
	UpdateAt    *gtime.Time `orm:"update_at"    json:"UpdatedAt"`   // 更新时间
	DeleteAt    *gtime.Time `orm:"delete_at"    json:"DeletedAt"`   // 删除时间
	Uuid        string      `orm:"uuid"         json:"uuid"`        // 用户唯一标识UUID
	Avatar      string      `orm:"avatar"       json:"headerImg"`   // 用户头像
	Nickname    string      `orm:"nickname"     json:"nickName"`    // 用户昵称
	Username    string      `orm:"username"     json:"userName"`    // 用户名
	Password    string      `orm:"password"     json:"-"`           // 用户登录密码
	AuthorityId string      `orm:"authority_id" json:"authorityId"` // 用户角色ID
}

//@author: SliverHorn
//@description: 密码检查(工具类)
//@return: bool(false 校验失败, true 校验成功)
func (a *Admin) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

//@author: SliverHorn
//@description: 加密密码(工具类)
func (a *Admin) EncryptedPassword() error {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost); err != nil { // 加密密码
		return err
	} else {
		a.Password = string(byTes)
		return nil
	}
}