package response

var Message = map[Code]string{
	ServerBusy: "服务器忙ing...",

	SuccessStart:        "成功码开始!",
	SuccessAdd:          "添加成功!",
	SuccessFirst:        "获取数据成功!",
	SuccessCreated:      "创建成功!",
	SuccessUpdated:      "更新成功!",
	SuccessDeleted:      "删除成功!",
	SuccessGetList:      "获取列表数据成功!",
	SuccessOperation:    "操作成功!",
	SuccessBatchDeleted: "批量删除成功!",
	SuccessEnd:          "成功码结束!",

	ErrorStart:        "失败码开始!",
	ErrorAdd:          "添加失败!",
	ErrorFirst:        "获取数据失败!",
	ErrorCreated:      "创建失败!",
	ErrorUpdated:      "更新失败!",
	ErrorDeleted:      "删除失败!",
	ErrorGetList:      "获取列表数据失败!",
	ErrorOperation:    "操作失败!",
	ErrorBatchDeleted: "批量删除失败!",
	ErrorEnd:          "失败码结束!",
}
