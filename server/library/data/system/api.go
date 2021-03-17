package data

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: apis 表数据初始化
func (a *api) Init() error {
	apis := []model.Api{
		{global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", I18nHash["UserLogin"], "base", "POST"},
		{global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/register", I18nHash["UserRegister"], "user", "POST"},
		{global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/createApi", I18nHash["CreateApi"], "api", "POST"},
		{global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiList", I18nHash["GetApiList"], "api", "POST"},
		{global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiById", I18nHash["GetApiDetail"], "api", "POST"},
		{global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApi", I18nHash["DeleteApi"], "api", "POST"},
		{global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/updateApi", I18nHash["UpdateApi"], "api", "POST"},
		{global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getAllApis", I18nHash["GetAllApis"], "api", "POST"},
		{global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/createAuthority", I18nHash["CreateAuthority"], "authority", "POST"},
		{global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/deleteAuthority", I18nHash["DeleteAuthority"], "authority", "POST"},
		{global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/getAuthorityList", I18nHash["GetAuthorityList"], "authority", "POST"},
		{global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenu", I18nHash["GetMenu"], "menu", "POST"},
		{global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuList", I18nHash["GetMenuList"], "menu", "POST"},
		{global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addBaseMenu", I18nHash["AddBaseMenu"], "menu", "POST"},
		{global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuTree", I18nHash["GetBaseMenuTree"], "menu", "POST"},
		{global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addMenuAuthority", I18nHash["AddMenuAuthority"], "menu", "POST"},
		{global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuAuthority", I18nHash["GetMenuAuthority"], "menu", "POST"},
		{global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/deleteBaseMenu", I18nHash["DeleteBaseMenu"], "menu", "POST"},
		{global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/updateBaseMenu", I18nHash["UpdateBaseMenu"], "menu", "POST"},
		{global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuById", I18nHash["GetBaseMenuById"], "menu", "POST"},
		{global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/changePassword", I18nHash["ChangePassword"], "user", "POST"},
		{global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/getUserList", I18nHash["GetUserList"], "user", "POST"},
		{global.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserAuthority", I18nHash["SetUserAuthority"], "user", "POST"},
		{global.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/upload", I18nHash["UploadFile"], "fileUploadAndDownload", "POST"},
		{global.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/getFileList", I18nHash["GetFileList"], "fileUploadAndDownload", "POST"},
		{global.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/updateCasbin", I18nHash["UpdateCasbin"], "casbin", "POST"},
		{global.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/getPolicyPathByAuthorityId", I18nHash["GetPolicyPathByAuthorityId"], "casbin", "POST"},
		{global.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/deleteFile", I18nHash["DeleteFile"], "fileUploadAndDownload", "POST"},
		{global.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/jwt/jsonInBlacklist", I18nHash["JsonInBlacklist"], "jwt", "POST"},
		{global.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setDataAuthority", I18nHash["SetDataAuthority"], "authority", "POST"},
		{global.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getSystemConfig", I18nHash["GetSystemConfig"], "system", "POST"},
		{global.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/setSystemConfig", I18nHash["SetSystemConfig"], "system", "POST"},
		{global.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", I18nHash["CreateCustomer"], "customer", "POST"},
		{global.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", I18nHash["UpdateCustomer"], "customer", "PUT"},
		{global.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", I18nHash["DeleteCustomer"], "customer", "DELETE"},
		{global.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", I18nHash["GetCustomer"], "customer", "GET"},
		{global.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customerList", I18nHash["GetCustomerList"], "customer", "GET"},
		{global.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/casbinTest/:pathParam", I18nHash["RESTFULTest"], "casbin", "GET"},
		{global.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/createTemp", I18nHash["CreateTemp"], "autoCode", "POST"},
		{global.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/updateAuthority", I18nHash["UpdateAuthority"], "authority", "PUT"},
		{global.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/copyAuthority", I18nHash["CopyAuthority"], "authority", "POST"},
		{global.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/deleteUser", I18nHash["DeleteUser"], "user", "DELETE"},
		{global.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/createSysDictionaryDetail", I18nHash["CreateSysDictionaryDetail"], "sysDictionaryDetail", "POST"},
		{global.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/deleteSysDictionaryDetail", I18nHash["DeleteSysDictionaryDetail"], "sysDictionaryDetail", "DELETE"},
		{global.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/updateSysDictionaryDetail", I18nHash["UpdateSysDictionaryDetail"], "sysDictionaryDetail", "PUT"},
		{global.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/findSysDictionaryDetail", I18nHash["FindSysDictionaryDetail"], "sysDictionaryDetail", "GET"},
		{global.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/getSysDictionaryDetailList", I18nHash["GetSysDictionaryDetailList"], "sysDictionaryDetail", "GET"},
		{global.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/createSysDictionary", I18nHash["CreateSysDictionary"], "sysDictionary", "POST"},
		{global.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/deleteSysDictionary", I18nHash["DeleteSysDictionary"], "sysDictionary", "DELETE"},
		{global.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/updateSysDictionary", I18nHash["UpdateSysDictionary"], "sysDictionary", "PUT"},
		{global.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/findSysDictionary", I18nHash["FindSysDictionary"], "sysDictionary", "GET"},
		{global.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/getSysDictionaryList", I18nHash["GetSysDictionaryList"], "sysDictionary", "GET"},
		{global.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/createSysOperationRecord", I18nHash["CreateSysOperationRecord"], "sysOperationRecord", "POST"},
		{global.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecord", I18nHash["DeleteSysOperationRecord"], "sysOperationRecord", "DELETE"},
		{global.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/findSysOperationRecord", I18nHash["FindSysOperationRecord"], "sysOperationRecord", "GET"},
		{global.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/getSysOperationRecordList", I18nHash["GetSysOperationRecordList"], "sysOperationRecord", "GET"},
		{global.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getTables", I18nHash["GetTables"], "autoCode", "GET"},
		{global.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getDB", I18nHash["GetDB"], "autoCode", "GET"},
		{global.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getColumn", I18nHash["GetColumn"], "autoCode", "GET"},
		{global.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecordByIds", I18nHash["DeleteSysOperationRecordByIds"], "sysOperationRecord", "DELETE"},
		{global.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/upload", I18nHash["SubsectionUpload"], "simpleUploader", "POST"},
		{global.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/checkFileMd5", I18nHash["CheckFileMd5"], "simpleUploader", "GET"},
		{global.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/mergeFileMd5", I18nHash["MergeFileMd5"], "simpleUploader", "GET"},
		{global.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserInfo", I18nHash["SetUserInfo"], "user", "PUT"},
		{global.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getServerInfo", I18nHash["GetServerInfo"], "system", "POST"},
		{global.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/email/emailTest", I18nHash["EmailTest"], "email", "POST"},
		{global.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/createWorkflowProcess", I18nHash["CreateWorkflowProcess"], "workflowProcess", "POST"},
		{global.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcess", I18nHash["DeleteWorkflowProcess"], "workflowProcess", "DELETE"},
		{global.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcessByIds", I18nHash["DeleteWorkflowProcessByIds"], "workflowProcess", "DELETE"},
		{global.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/updateWorkflowProcess", I18nHash["UpdateWorkflowProcess"], "workflowProcess", "PUT"},
		{global.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowProcess", I18nHash["FindWorkflowProcess"], "workflowProcess", "GET"},
		{global.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowProcessList", I18nHash["GetWorkflowProcessList"], "workflowProcess", "GET"},
		{global.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowStep", I18nHash["FindWorkflowStep"], "workflowProcess", "GET"},
		{global.Model{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/startWorkflow", I18nHash["StartWorkflow"], "workflowProcess", "POST"},
		{global.Model{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyStated", I18nHash["GetMyStated"], "workflowProcess", "GET"},
		{global.Model{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyNeed", I18nHash["GetMyNeed"], "workflowProcess", "GET"},
		{global.Model{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowMoveByID", I18nHash["GetWorkflowMoveByID"], "workflowProcess", "GET"},
		{global.Model{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/completeWorkflowMove", I18nHash["CompleteWorkflowMove"], "workflowProcess", "POST"},
		{global.Model{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/preview", I18nHash["Preview"], "autoCode", "POST"},
		{global.Model{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/importExcel", I18nHash["ImportExcel"], "excel", "POST"},
		{global.Model{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/loadExcel", I18nHash["LoadExcel"], "excel", "GET"},
		{global.Model{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/exportExcel", I18nHash["ExportExcel"], "excel", "POST"},
		{global.Model{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/downloadTemplate", I18nHash["DownloadTemplate"], "excel", "GET"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 84}).Find(&[]model.Api{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (a *api) TableName() string {
	return "apis"
}
