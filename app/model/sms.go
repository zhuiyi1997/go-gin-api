package model

type Sms struct {
	ID uint64 `json:"id" form:"id" gorm:"primary_key"`
	Phone string `json:"phone" form:"phone" gorm:"type:varchar(20);not null"`
	SmsCode string `json:"sms_code" form:"sms_code" gorm:"type:varchar(10);not null"`
	Type string `json:"type" form:"type" gorm:"type:varchar(30);not null"`
	SendTime string `json:"send_time" form:"send_time" gorm:"type:timestamp"`
}

func (sms *Sms) TableName() string {
	return "sms"
}
