package response

type Customer struct {
	CustomerName 		string `json:"customerName" gorm:"comment:客户名"`			//客户名
	CustomerPhoneData 	string `json:"customerPhoneData" gorm:"comment:客户手机号码"`	//客户手机号码
	SysUserID 			uint `json:"sysUserId" gorm:"comment:管理ID"`				//管理id
	SysUserAuthorityID 	string `json:"sysUserAuthorityID" gorm:"comment:管理角色ID"` //管理角色id

}