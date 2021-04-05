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
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/base/login", Description: I18nHash["UserLogin"], ApiGroup: "base", Method: "POST"},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/register", Description: I18nHash["UserRegister"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/createApi", Description: I18nHash["CreateApi"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiList", Description: I18nHash["GetApiList"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiById", Description: I18nHash["GetApiDetail"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApi", Description: I18nHash["DeleteApi"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/updateApi", Description: I18nHash["UpdateApi"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getAllApis", Description: I18nHash["GetAllApis"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/createAuthority", Description: I18nHash["CreateAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/deleteAuthority", Description: I18nHash["DeleteAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/getAuthorityList", Description: I18nHash["GetAuthorityList"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenu", Description: I18nHash["GetMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuList", Description: I18nHash["GetMenuList"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addBaseMenu", Description: I18nHash["AddBaseMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuTree", Description: I18nHash["GetBaseMenuTree"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addMenuAuthority", Description: I18nHash["AddMenuAuthority"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuAuthority", Description: I18nHash["GetMenuAuthority"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/deleteBaseMenu", Description: I18nHash["DeleteBaseMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/updateBaseMenu", Description: I18nHash["UpdateBaseMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuById", Description: I18nHash["GetBaseMenuById"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/changePassword", Description: I18nHash["ChangePassword"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getUserList", Description: I18nHash["GetUserList"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserAuthority", Description: I18nHash["SetUserAuthority"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/upload", Description: I18nHash["UploadFile"], ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/getFileList", Description: I18nHash["GetFileList"], ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/updateCasbin", Description: I18nHash["UpdateCasbin"], ApiGroup: "casbin", Method: "POST"},
		{Model: global.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/getPolicyPathByAuthorityId", Description: I18nHash["GetPolicyPathByAuthorityId"], ApiGroup: "casbin", Method: "POST"},
		{Model: global.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/deleteFile", Description: I18nHash["DeleteFile"], ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/jwt/jsonInBlacklist", Description: I18nHash["JsonInBlacklist"], ApiGroup: "jwt", Method: "POST"},
		{Model: global.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/setDataAuthority", Description: I18nHash["SetDataAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getSystemConfig", Description: I18nHash["GetSystemConfig"], ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/setSystemConfig", Description: I18nHash["SetSystemConfig"], ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: I18nHash["CreateCustomer"], ApiGroup: "customer", Method: "POST"},
		{Model: global.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: I18nHash["UpdateCustomer"], ApiGroup: "customer", Method: "PUT"},
		{Model: global.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: I18nHash["DeleteCustomer"], ApiGroup: "customer", Method: "DELETE"},
		{Model: global.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: I18nHash["GetCustomer"], ApiGroup: "customer", Method: "GET"},
		{Model: global.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customerList", Description: I18nHash["GetCustomerList"], ApiGroup: "customer", Method: "GET"},
		{Model: global.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/casbinTest/:pathParam", Description: I18nHash["RESTFULTest"], ApiGroup: "casbin", Method: "GET"},
		{Model: global.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/createTemp", Description: I18nHash["CreateTemp"], ApiGroup: "autoCode", Method: "POST"},
		{Model: global.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/updateAuthority", Description: I18nHash["UpdateAuthority"], ApiGroup: "authority", Method: "PUT"},
		{Model: global.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/copyAuthority", Description: I18nHash["CopyAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/deleteUser", Description: I18nHash["DeleteUser"], ApiGroup: "user", Method: "DELETE"},
		{Model: global.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: I18nHash["CreateSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "POST"},
		{Model: global.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: I18nHash["DeleteSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "DELETE"},
		{Model: global.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: I18nHash["UpdateSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "PUT"},
		{Model: global.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: I18nHash["FindSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: global.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: I18nHash["GetSysDictionaryDetailList"], ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: global.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/createSysDictionary", Description: I18nHash["CreateSysDictionary"], ApiGroup: "sysDictionary", Method: "POST"},
		{Model: global.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/deleteSysDictionary", Description: I18nHash["DeleteSysDictionary"], ApiGroup: "sysDictionary", Method: "DELETE"},
		{Model: global.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/updateSysDictionary", Description: I18nHash["UpdateSysDictionary"], ApiGroup: "sysDictionary", Method: "PUT"},
		{Model: global.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/findSysDictionary", Description: I18nHash["FindSysDictionary"], ApiGroup: "sysDictionary", Method: "GET"},
		{Model: global.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/getSysDictionaryList", Description: I18nHash["GetSysDictionaryList"], ApiGroup: "sysDictionary", Method: "GET"},
		{Model: global.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/createSysOperationRecord", Description: I18nHash["CreateSysOperationRecord"], ApiGroup: "sysOperationRecord", Method: "POST"},
		{Model: global.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecord", Description: I18nHash["DeleteSysOperationRecord"], ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: global.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/findSysOperationRecord", Description: I18nHash["FindSysOperationRecord"], ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: global.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/getSysOperationRecordList", Description: I18nHash["GetSysOperationRecordList"], ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: global.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getTables", Description: I18nHash["GetTables"], ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getDB", Description: I18nHash["GetDB"], ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getColumn", Description: I18nHash["GetColumn"], ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: I18nHash["DeleteSysOperationRecordByIds"], ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: global.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/upload", Description: I18nHash["SubsectionUpload"], ApiGroup: "simpleUploader", Method: "POST"},
		{Model: global.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/checkFileMd5", Description: I18nHash["CheckFileMd5"], ApiGroup: "simpleUploader", Method: "GET"},
		{Model: global.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/mergeFileMd5", Description: I18nHash["MergeFileMd5"], ApiGroup: "simpleUploader", Method: "GET"},
		{Model: global.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserInfo", Description: I18nHash["SetUserInfo"], ApiGroup: "user", Method: "PUT"},
		{Model: global.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getServerInfo", Description: I18nHash["GetServerInfo"], ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/email/emailTest", Description: I18nHash["EmailTest"], ApiGroup: "email", Method: "POST"},
		{Model: global.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/createWorkflowProcess", Description: I18nHash["CreateWorkflowProcess"], ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/deleteWorkflowProcess", Description: I18nHash["DeleteWorkflowProcess"], ApiGroup: "workflowProcess", Method: "DELETE"},
		{Model: global.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/deleteWorkflowProcessByIds", Description: I18nHash["DeleteWorkflowProcessByIds"], ApiGroup: "workflowProcess", Method: "DELETE"},
		{Model: global.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/updateWorkflowProcess", Description: I18nHash["UpdateWorkflowProcess"], ApiGroup: "workflowProcess", Method: "PUT"},
		{Model: global.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/findWorkflowProcess", Description: I18nHash["FindWorkflowProcess"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getWorkflowProcessList", Description: I18nHash["GetWorkflowProcessList"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/findWorkflowStep", Description: I18nHash["FindWorkflowStep"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/startWorkflow", Description: I18nHash["StartWorkflow"], ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getMyStated", Description: I18nHash["GetMyStated"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getMyNeed", Description: I18nHash["GetMyNeed"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getWorkflowMoveByID", Description: I18nHash["GetWorkflowMoveByID"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/completeWorkflowMove", Description: I18nHash["CompleteWorkflowMove"], ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/preview", Description: I18nHash["Preview"], ApiGroup: "autoCode", Method: "POST"},
		{Model: global.Model{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/importExcel", Description: I18nHash["ImportExcel"], ApiGroup: "excel", Method: "POST"},
		{Model: global.Model{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/loadExcel", Description: I18nHash["LoadExcel"], ApiGroup: "excel", Method: "GET"},
		{Model: global.Model{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/exportExcel", Description: I18nHash["ExportExcel"], ApiGroup: "excel", Method: "POST"},
		{Model: global.Model{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/downloadTemplate", Description: I18nHash["DownloadTemplate"], ApiGroup: "excel", Method: "GET"},
		{Model: global.Model{ID: 85, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApisByIds", Method: "批量删除api", ApiGroup: "api", Description: "DELETE"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 85}).Find(&[]model.Api{}).RowsAffected == 2 {
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
