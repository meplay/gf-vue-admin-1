package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRegister struct {
	Avatar       string   `json:"headerImg"`
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	Nickname     string   `json:"nickName"`
	AuthorityId  string   `json:"authorityId"`
	AuthorityIds []string `json:"authorityIds"`
}

func (r *UserRegister) Create() system.User {
	length := len(r.AuthorityIds)
	authorities := make([]system.Authority, 0, length)
	for i := 0; i < length; i++ {
		authorities = append(authorities, system.Authority{AuthorityId: r.AuthorityIds[i]})
	}
	return system.User{
		Uuid:        uuid.NewV4().String(),
		Avatar:      r.Avatar,
		Username:    r.Username,
		Password:    r.Password,
		Nickname:    r.Nickname,
		AuthorityId: r.AuthorityId,
		Authorities: authorities,
	}
}

type UserLogin struct {
	Captcha   string `json:"captcha" example:"验证码"`
	Username  string `json:"username" example:"用户名"`
	Password  string `json:"password" example:"密码"`
	CaptchaId string `json:"captchaId" example:"验证码id"`
}

type UserFind struct {
	ID   uint   `json:"id" example:"7"`
	Uuid string `json:"uuid" example:"uuid"`
}

func (r *UserFind) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.ID != 0 {
			db = db.Where("id = ?", r.ID)
		}
		if r.Uuid != "" {
			db = db.Where("uuid = ?", r.Uuid)
		}
		return db
	}
}

type UserUpdate struct {
	common.GetByID
	Avatar      string `json:"headerImg" example:"用户头像"`
	Username    string `json:"userName" example:"用户登录名"`
	Nickname    string `json:"nickName" example:"用户昵称"`
	SideMode    string `json:"sideMode" example:"用户侧边主题"`
	BaseColor   string `json:"baseColor" example:"基础颜色"`
	ActiveColor string `json:"activeColor" example:"用户角色ID"`
}

func (r *UserUpdate) Update() system.User {
	return system.User{
		Avatar:      r.Avatar,
		Username:    r.Username,
		Nickname:    r.Nickname,
		SideMode:    r.SideMode,
		BaseColor:   r.BaseColor,
		ActiveColor: r.ActiveColor,
	}
}

type UserChangePassword struct {
	Uuid        string
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type UserSetAuthority struct {
	ID          uint   `json:"-"`
	Uuid        string `json:"-"`
	AuthorityId string `json:"authorityId"`
}

type UserSetAuthorities struct {
	ID           uint     `json:"ID" example:"7"`
	AuthorityIds []string `json:"authorityIds" example:"角色id切片"`
}
