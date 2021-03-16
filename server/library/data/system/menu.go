package data

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"time"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var (
	Menu = new(menu)
)

type menu struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: menus 表数据初始化
func (m *menu) Init() error {
	var menus = []model.Menu{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: I18nHash["InstrumentPanel"], Icon: "setting"}},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: model.Meta{Title: I18nHash["AboutUs"], Icon: "info"}},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: I18nHash["SuperAdmin"], Icon: "user-solid"}},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: I18nHash["RoleManagement"], Icon: "s-custom"}},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: I18nHash["MenuManagement"], Icon: "s-order", KeepAlive: true}},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: I18nHash["APIManagement"], Icon: "s-platform", KeepAlive: true}},
		{Model: global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: model.Meta{Title: I18nHash["UserManagement"], Icon: "coordinate"}},
		{Model: global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: I18nHash["PersonalInformation"], Icon: "message-solid"}},
		{Model: global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: model.Meta{Title: I18nHash["SampleFile"], Icon: "s-management"}},
		{Model: global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: model.Meta{Title: I18nHash["ExcelImportExport"], Icon: "s-marketing"}},
		{Model: global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: model.Meta{Title: I18nHash["MediaLibrary"], Icon: "upload"}},
		{Model: global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: model.Meta{Title: I18nHash["BreakpointContinuation"], Icon: "upload"}},
		{Model: global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: model.Meta{Title: I18nHash["CustomerList"], Icon: "s-custom"}},
		{Model: global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: I18nHash["SystemTool"], Icon: "s-cooperation"}},
		{Model: global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: I18nHash["CodeGenerator"], Icon: "cpu", KeepAlive: true}},
		{Model: global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: I18nHash["FormGenerator"], Icon: "magic-stick", KeepAlive: true}},
		{Model: global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: I18nHash["SystemConfiguration"], Icon: "s-operation"}},
		{Model: global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: I18nHash["DictionaryManagement"], Icon: "notebook-2"}},
		{Model: global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: I18nHash["DictionaryDetails"], Icon: "s-order"}},
		{Model: global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: I18nHash["OperationHistory"], Icon: "time"}},
		{Model: global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: model.Meta{Title: I18nHash["BreakpointContinuationPlug"], Icon: "upload"}},
		{Model: global.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Hidden: false, Component: "/", Sort: 0, Meta: model.Meta{Title: I18nHash["OfficialWebsite"], Icon: "s-home"}},
		{Model: global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: I18nHash["ServerStatus"], Icon: "cloudy"}},
		{Model: global.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "workflow", Name: "workflow", Hidden: false, Component: "view/workflow/index.vue", Sort: 5, Meta: model.Meta{Title: I18nHash["WorkflowFunction"], Icon: "phone"}},
		{Model: global.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowCreate", Name: "workflowCreate", Hidden: false, Component: "view/workflow/workflowCreate/workflowCreate.vue", Sort: 0, Meta: model.Meta{Title: I18nHash["WorkflowDrawing"], Icon: "circle-plus"}},
		{Model: global.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowProcess", Name: "workflowProcess", Hidden: false, Component: "view/workflow/workflowProcess/workflowProcess.vue", Sort: 0, Meta: model.Meta{Title: I18nHash["WorkflowList"], Icon: "s-cooperation"}},
		{Model: global.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowUse", Name: "workflowUse", Hidden: true, Component: "view/workflow/workflowUse/workflowUse.vue", Sort: 0, Meta: model.Meta{Title: I18nHash["UseWorkflow"], Icon: "video-play"}},
		{Model: global.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "started", Name: "started", Hidden: false, Component: "view/workflow/userList/started.vue", Sort: 0, Meta: model.Meta{Title: I18nHash["IInitiatedIt"], Icon: "s-order"}},
		{Model: global.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "need", Name: "need", Hidden: false, Component: "view/workflow/userList/need.vue", Sort: 0, Meta: model.Meta{Title: I18nHash["MyToDoList"], Icon: "s-platform"}},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]model.Menu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (m *menu) TableName() string {
	return "menus"
}
