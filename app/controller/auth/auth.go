package auth

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/zhuiyi1997/go-gin-api/app/controller/param_bind"
	"github.com/zhuiyi1997/go-gin-api/app/controller/param_verify"
	"github.com/zhuiyi1997/go-gin-api/app/model"
	"github.com/zhuiyi1997/go-gin-api/app/util/bind"
	"github.com/zhuiyi1997/go-gin-api/app/util/response"
	"github.com/zhuiyi1997/go-gin-api/app/util/sms"
	"github.com/zhuiyi1997/go-gin-api/app/util/iot"
	"github.com/zhuiyi1997/go-gin-api/app/util/validator_trans"
	"gopkg.in/go-playground/validator.v9"
	"crypto/sha1"
	"log"
	_ "log"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func Login(c *gin.Context)  {
	utilGin := response.Gin{Ctx: c}

	s, e := bind.Bind(&param_bind.Login{}, c)
	if e != nil {
		utilGin.Response(-1, e.Error(), nil)
		return
	}

	// 参数验证
	validate := validator.New()
	trans := validator_trans.SetZh(validate)

	// 注册自定义验证
	_ = validate.RegisterValidation("CheckPhone", param_verify.CheckPhone)
	_ = validate.RegisterValidation("CheckSms", param_verify.CheckSms)

	if err := validate.Struct(s); err != nil {
		for _,one := range err.(validator.ValidationErrors).Translate(trans) {
			log.Println(one)
			utilGin.Response(-1, one, nil)
			return
		}
	}
	// 业务处理...
	// 1.是否存在用户
	var user model.RegUser
	var loginApi param_bind.LoginApi
	mysql := model.NewGorm()
	local_time := time.Now().Unix()
	local_datetime := time.Now().Format("2006-01-02 15:04:05")
	sh := sha1.New()
	sh.Write([]byte(c.PostForm("phone")+strconv.FormatInt(local_time,10)+strconv.Itoa(rand.Intn(1000))))
	sign := hex.EncodeToString(sh.Sum(nil))
	mysql.Db.Select("id,name,phone,is_vip,IotDeviceName,IotDeviceSecret,AliIotId").Where("phone = ?",c.PostForm("phone")).First(&user)
	if user.ID == 0 {
		//注册阿里云iot
		d,_ := iot.NewIot()
		var device = map[string]string{"DeviceName":c.PostForm("phone"),"Nickname":c.PostForm("phone")}
		b,result := d.RegisterDevice(device)
		log.Println(b,result)
		if b {
			log.Println(hex.EncodeToString(sh.Sum(nil)))
			var u = model.RegUser{Name:c.PostForm("phone"),Phone:c.PostForm("phone"),LoginSign: sign,
				LastLoginTime: int32(local_time),IotDeviceName: result["DeviceName"],IotDeviceSecret: result["DeviceSecret"],
			AliIotId: result["IotId"],CreatedAt: local_datetime,UpdatedAt: local_datetime}
			mysql.Db.Create(&u)
			mysql.Db.Table("reg_user").Where("phone = ?",c.PostForm("phone")).Scan(&loginApi)
			utilGin.Response(1, "success", loginApi)
		} else {
			utilGin.Response(-1, result["error"], nil)
			return
		}
	} else {
		if user.IsLocked == '1' {
			utilGin.Response(-1, "账号已被锁定", nil)
			return
		}
		// 更新
		mysql.Db.Table("reg_user").Where("phone = ?",c.PostForm("phone")).Updates(map[string]interface{}{"login_sign": sign,"updated_at":local_datetime,
			"last_login_time": int32(local_time),})
		mysql.Db.Table("reg_user").Where("phone = ?",c.PostForm("phone")).Scan(&loginApi)
		utilGin.Response(1, "success", loginApi)
	}
	defer func() {
		mysql.Db.Close()
	}()
}

// 发送验证码
func SendSms(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	s, e := bind.Bind(&param_bind.SendSms{}, c)
	if e != nil {
		utilGin.Response(-1, e.Error(), nil)
		return
	}

	// 参数验证
	validate := validator.New()
	trans := validator_trans.SetZh(validate)

	// 注册自定义验证
	_ = validate.RegisterValidation("CheckPhone", param_verify.CheckPhone)

	if err := validate.Struct(s); err != nil {
		for _,one := range err.(validator.ValidationErrors).Translate(trans) {
			log.Println(one)
			utilGin.Response(-1, one, nil)
			return
		}
	}

	// 业务处理...
	field := reflect.ValueOf(s).Elem()
	_, err := sms.SendSms(field.FieldByName("Type").String(),field.FieldByName("Phone").String())
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	utilGin.Response(1, "success", nil)
}
