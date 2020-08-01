package param_bind

// 登录接口参数
type Login struct {
	Phone string `json:"phone" form:"phone" validate:"required,CheckPhone"`
	SmsCode string `json:"sms_code" form:"sms_code" validate:"required,CheckSms"`
}

//　登录返回参数
type LoginApi struct {
	ID uint64 `json:"id" gorm:"column:id"`
	SessionKey string `json:"session_key" gorm:"column:login_sign"`
	UserName string `json:"userName" gorm:"column:name"`
	Phone string `json:"phone" gorm:"column:phone"`
	DeviceName string `json:"IotDeviceName" gorm:"column:IotDeviceName"`
	DeviceSecret string `json:"DeviceSecret" gorm:"column:IotDeviceSecret"`
	IotId string `json:"IotId" gorm:"column:AliIotId"`
}

// 发送验证码参数
type SendSms struct {
	Type string `json:"type" form:"type" validate:"required,oneof=login logoff"`
	Phone string `json:"phone" form:"phone" validate:"required,CheckPhone"`
}
