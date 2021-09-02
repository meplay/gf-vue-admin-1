package response

type Code int

const (
	// ServerBusy 基础码
	ServerBusy Code = iota

	// SuccessStart 成功开始标记
	SuccessStart

	SuccessAdd
	SuccessFirst
	SuccessCreated
	SuccessDeleted
	SuccessUpdated
	SuccessGetList
	SuccessOperation
	SuccessBatchDeleted

	// SuccessStart 成功结束标记
	SuccessEnd

	// ErrorStart 失败开始标记
	ErrorStart

	ErrorAdd
	ErrorFirst
	ErrorCreated
	ErrorUpdated
	ErrorDeleted
	ErrorGetList
	ErrorOperation
	ErrorBatchDeleted

	// ErrorEnd 失败结束标记
	ErrorEnd
)

// Message code 对应 massage
// Author [SliverHorn](https://github.com/SliverHorn)
func (c Code) Message() string {
	if msg, ok := Message[c]; ok {
		return msg
	}
	return Message[ServerBusy]
}
