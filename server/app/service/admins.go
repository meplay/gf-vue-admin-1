package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/admins"
	"server/library/global"
	"server/library/utils"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
)

// ChangePassword Change administrator password
// ChangePassword 修改管理员密码
func ChangePassword(change *request.ChangePassword) (err error) {
	var admin *admins.Entity
	if admin, err = admins.FindOne(g.Map{"username": change.Username}); err != nil {
		return errors.New("用户不存在, 修改密码失败")
	}
	if utils.CompareHashAndPassword(admin.Password, change.Password) {
		if admin.Password, err = utils.EncryptedPassword(change.NewPassword); err != nil {
			return errors.New("修改密码失败")
		}
		if _, err = admins.Save(admin); err != nil {
			return errors.New("修改密码失败")
		}
		return
	}
	return errors.New("旧密码输入有误")
}

// UploadHeaderImg User uploads profile picture
// UploadHeaderImg 用户上传头像
func UploadHeaderImg(userUuid string, filePath string) (adminInfo *admins.Entity, err error) {
	if _, err := admins.Update(g.Map{"header_img": filePath}, g.Map{"uuid": userUuid}); err != nil {
		return adminInfo, errors.New("")
	}
	adminInfo, err = admins.FindOne(g.Map{"uuid": userUuid})
	return adminInfo, err
}

// GetAdminList Paging gets the list of users
// GetAdminList 分页获取用户列表
func GetAdminList(info *request.PageInfo) (list interface{}, total int, err error) {
	adminList := ([]*admins.AdminHasOneAuthority)(nil)
	adminDb := g.DB(global.Db).Table("admins").Safe()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	total, err = adminDb.Count()
	err = adminDb.Limit(limit).Offset(offset).ScanList(&adminList, "Admin")
	err = adminDb.Where("authority_id", gdb.ListItemValues(adminList, "Admin", "AuthorityId")).ScanList(&adminList, "Authority", "Admin", "authority_id:AuthorityId")
	return adminList, total, err
}

// FindAdmin Used to refresh token and return admin information according to uuid
// FindAdmin 用于刷新token,根据uuid返回admin信息
func FindAdmin(adminUUID string) (admin *admins.AdminHasOneAuthority, err error) {
	admin = (*admins.AdminHasOneAuthority)(nil)
	db := g.DB(global.Db).Table("admins").Safe()
	err = db.Where(g.Map{"uuid": adminUUID}).Struct(&admin)
	err = db.Struct(&admin.Authority, g.Map{"authority_id": admin.AuthorityId})
	return admin, err
}

// SetUserAuthority Set user permissions
// SetUserAuthority 设置用户权限
func SetUserAuthority(set *request.SetAdminAuthority) (err error) {
	_, err = admins.Update(g.Map{"authority_id": set.AuthorityId}, g.Map{"uuid": set.Uuid})
	return err
}

// DeleteAdmin Delete administrator
// DeleteAdmin 删除管理员
func DeleteAdmin(delete *request.DeleteById) (err error) {
	_, err = admins.Delete(g.Map{"id": delete.Id})
	return err
}

// SetAdminInfo Set admin information
// SetAdminInfo 设置管理员信息
func SetAdminInfo(set *request.SetAdminInfo) (admin *admins.AdminHasOneAuthority, err error) {
	_, err = admins.Update(g.Map{"header_img": set.HeaderImg}, g.Map{"uuid": set.Uuid})
	admin, err = FindAdmin(set.Uuid)
	return admin, err
}
