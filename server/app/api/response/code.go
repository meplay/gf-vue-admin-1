package response

type Code int

const (
	// 基础码
	ServerBusy Code = iota

	// 成功
	SuccessStart

	SuccessAdd
	SuccessFirst
	SuccessCreated
	SuccessDeleted
	SuccessUpdated
	SuccessGetList
	SuccessOperation
	SuccessBatchDeleted

	// admin
	SuccessAdminLogin
	SuccessSetAuthority
	SuccessSetAdminInfo
	SuccessAdminRegister
	SuccessChangePassword

	// captcha
	SuccessCaptcha

	// Authority
	SuccessCopyAuthority
	SuccessCreateAuthority
	SuccessSetDataAuthority

	// JwtBlackList
	SuccessJwtBlackList

	// BreakpointContinue
	SuccessFind
	SuccessFinish
	SuccessCreateChunk
	SuccessRemoveChunk

	SuccessEnd

	//失败
	ErrorStart

	ErrorAdd
	ErrorFirst
	ErrorCreated
	ErrorUpdated
	ErrorDeleted
	ErrorGetList
	ErrorOperation
	ErrorBatchDeleted

	// admin
	ErrorAdminLogin
	ErrorSetAuthority
	ErrorSetAdminInfo
	ErrorAdminRegister
	ErrorChangePassword

	// captcha
	ErrorCaptcha

	// Authority
	ErrorCopyAuthority
	ErrorCreateAuthority
	ErrorSetDataAuthority

	// JwtBlackList
	ErrorJwtBlackList

	// BreakpointContinue
	ErrorFind
	ErrorFinish
	ErrorFormFile
	ErrorCreateChunk
	ErrorRemoveChunk

	ErrorEnd

	// 自定义码
	AdminNotFind
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: code 对应 massage
func (c Code) Message() string {
	if msg, ok := Message[c]; ok {
		return msg
	}
	return Message[ServerBusy]
}