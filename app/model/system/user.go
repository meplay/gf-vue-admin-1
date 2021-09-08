package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	global.Model
	Uuid        string `json:"uuid" gorm:"comment:用户UUID"`                                                    // 用户UUID
	Avatar      string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	Username    string `json:"userName" gorm:"comment:用户登录名"`                                                 // 用户登录名
	Password    string `json:"-"  gorm:"comment:用户登录密码"`                                                      // 用户登录密码
	Nickname    string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                     // 用户昵称
	SideMode    string `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                   // 用户侧边主题
	BaseColor   string `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                    // 基础颜色
	ActiveColor string `json:"activeColor" gorm:"default:#1890ff;comment:用户角色ID"`                             // 活跃颜色
	AuthorityId string `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                 // 用户角色ID

	Authority   Authority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []Authority `json:"authorities" gorm:"many2many:users_authorities"`
}

func (u *User) TableName() string {
	return "users"
}

// CompareHashAndPassword 密码检查 false 校验失败, true 校验成功
// Author [SliverHorn](https://github.com/SliverHorn)
func (u *User) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

// EncryptedPassword 加密密码
// Author [SliverHorn](https://github.com/SliverHorn)
func (u *User) EncryptedPassword() error {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil { // 加密密码
		return err
	} else {
		u.Password = string(byTes)
		return nil
	}
}
