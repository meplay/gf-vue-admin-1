package response

var Message = map[Code]string{
	// 基础信息
	ServerBusy: "服务器忙ing...",

	// 成功
	SuccessStart: "成功码开始!",

	SuccessAdd:          "添加成功!",
	SuccessFirst:        "获取一条数据成功!",
	SuccessCreated:       "创建成功!",
	SuccessUpdated:       "更新成功!",
	SuccessDeleted:      "删除成功!",
	SuccessGetList:      "获取列表数据成功!",
	SuccessOperation:    "操作成功!",
	SuccessBatchDeleted: "批量删除成功!",

	// Admin
	SuccessAdminLogin:     "登录成功!",
	SuccessSetAuthority:   "设置角色成功!",
	SuccessSetAdminInfo:   "更新用户信息成功!",
	SuccessAdminRegister:  "注册成功!",
	SuccessChangePassword: "修改密码成功!",

	// Captcha
	SuccessCaptcha: "验证码获取成功!",

	// Authority
	SuccessCopyAuthority:    "复制角色成功!",
	SuccessCreateAuthority:  "创建角色成功!",
	SuccessSetDataAuthority: "创建角色成功!",

	// JwtBlackList
	SuccessJwtBlackList: "jwt作废成功!",

	SuccessEnd: "成功码结束!",

	//失败
	ErrorStart: "失败码开始!",

	ErrorAdd:          "添加失败!",
	ErrorFirst:        "获取一条数据失败!",
	ErrorCreated:       "创建失败!",
	ErrorUpdated:       "更新失败!",
	ErrorDeleted:      "删除失败!",
	ErrorGetList:      "获取列表数据失败!",
	ErrorOperation:    "操作失败!",
	ErrorBatchDeleted: "批量删除失败!",

	// Admin
	ErrorAdminLogin:     "登录失败!",
	ErrorSetAuthority:   "设置角色失败!",
	ErrorSetAdminInfo:   "更新用户信息失败!",
	ErrorAdminRegister:  "注册失败!",
	ErrorChangePassword: "修改密码失败!",

	// Captcha
	ErrorCaptcha: "验证码获取失败!",

	// Authority
	ErrorCopyAuthority:    "复制角色失败!",
	ErrorCreateAuthority:  "创建角色失败!",
	ErrorSetDataAuthority: "设置失败失败!",

	// JwtBlackList
	ErrorJwtBlackList: "jwt作废失败!",

	ErrorEnd: "失败码结束!",

	// 自定义信息
	AdminNotFind: "管理员用户不存在!",
}
