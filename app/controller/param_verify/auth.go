package param_verify

import (
	"github.com/zhuiyi1997/go-gin-api/app/model"
	"gopkg.in/go-playground/validator.v9"
	_ "log"
	"regexp"
	"time"
)

//　验证手机号
func CheckPhone(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if match, _ := regexp.MatchString(`^1[2-9]\d{9}$`,val);match{
		return true;
	}
	return false;
}

// 验证短信验证码
func CheckSms(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	mysql := model.NewGorm()
	local_time := time.Now().Unix()

	phone := fl.Top().Elem().FieldByName("Phone").String()
	var smsModel model.Sms
	mysql.Db.Where(&model.Sms{Phone:phone,Type:"login"}).Order("send_time desc").Attrs(model.Sms{SmsCode: "",SendTime: time.Now().Format("2006-01-02 15:04:05")}).FirstOrInit(&smsModel)
	if smsModel.SmsCode != "" && smsModel.SmsCode == val {
		loc, _ := time.LoadLocation("Local")
		tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", smsModel.SendTime, loc)

		if local_time - tmp.Unix() < 600 {
			return true;
		}
	}
	defer func() {
		mysql.Db.Close()
	}()
	return false;
}
