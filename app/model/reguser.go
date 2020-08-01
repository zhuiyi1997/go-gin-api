package model

import (
	_ "database/sql"
	_"time"
)

type RegUser struct {
	ID uint64 `json:"id,omitempty" form:"id" gorm:"type:bigint;unsigned;primary_key:true"`
	Name string `json:"userName,omitempty" form:"name" gorm:"type:varchar(50);column:name;"`  // 用户名
	Sex int8 `json:"sex,omitempty" form:"sex" gorm:"type:tinyint(1);default:3"`   //　性别
	Age int8 `json:"age,omitempty" form:"age" gorm:"type:tinyint(1)"`   //　年龄
	Avator string `json:"avator,omitempty" form:"avator" gorm:"type:varchar(255)"` //头像
	Phone string  `json:"phone,omitempty" form:"phone" gorm:"type:varchar(20);unique_index" binding:"required"` //账号
	LoginSign string `json:"session_key,omitempty" form:"login_sign" gorm:"type:varchar(255)"` //登录标志
	LastLoginTime int32 `json:"last_login_time,omitempty" form:"last_login_time" gorm:"type:int"` //最后登录时间
	Country string `json:"country,omitempty" form:"country" gorm:"type:varchar(50)"` //国家
	Province string `json:"province,omitempty" form:"province" gorm:"type:varchar(20)"` //省份
	City string `json:"city,omitempty" form:"city" gorm:"type:varchar(30)"` //城市
	Area string `json:"area,omitempty" form:"area" gorm:"type:varchar(255)"` //详细住址
	IsLocked int8 `json:"is_locked,omitempty" form:"is_locked" gorm:"type:tinyint(1);default:0;not null"` //账号是否有效
	IsVip int8 `json:"is_vip,omitempty" form:"is_vip" gorm:"type:tinyint(1);default:0;not null"` //是否是vip
	IotDeviceName string `json:"DeviceName,omitempty" gorm:"column:IotDeviceName;type:varchar(255);not null"` //物联网平台设备名称
	IotDeviceSecret string `json:"DeviceSecret,omitempty" gorm:"column:IotDeviceSecret;type:varchar(255);not null"` //物联网平台设备密钥
	AliIotId string `json:"IotId,omitempty" gorm:"column:AliIotId;type:varchar(255);not null"` //物联网平台设备唯一标识
	CreatedAt string `json:"created_at,omitempty" gorm:"type:timestamp"` //创建时间
	UpdatedAt string `json:"updated_at,omitempty" gorm:"type:timestamp"` //最后修改时间
}

func (reguser *RegUser) TableName() string {
	return "reg_user"
}
