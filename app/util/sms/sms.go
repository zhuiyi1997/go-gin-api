package sms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zhuiyi1997/go-gin-api/app/config"
	"github.com/zhuiyi1997/go-gin-api/app/model"
	"github.com/zhuiyi1997/go-gin-api/app/util/request"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

func SendSms(tp,phone string) (bool,error) {
	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(999999-100000)+100000)
	mysql := model.NewGorm()
	var sms_data model.Sms = model.Sms{Phone:phone,SmsCode: code,SendTime: time.Now().Format("2006-01-02 15:04:05"),Type: tp}

	var (
		content string
		boolean bool
		arr interface{}
	)
	boolean = true
	// 发送验证码
	if tp == "forget_password" {
		content = "【"+ config.SmsConf["sms_sign"] +"】您正在进行修改密码操作，您的验证码是："+ code +"，10分钟后失效，请及时验证。";
	} else {
		content = "【"+ config.SmsConf["sms_sign"] +"】您的验证码是："+ code +"，10分钟内有效，请勿告诉他人！";
	}
	escapeContent := url.QueryEscape(content)
	gateway := "http://sh2.ipyy.com/smsJson.aspx?action=send&userid="+config.SmsConf["userid"]+"&account="+config.SmsConf["account"]+"&password="+config.SmsConf["password"]+"&mobile="+phone+"&content="+escapeContent+"&sendTime=";
	log.Println(gateway)

	resp,err := request.HttpGet(gateway,&gin.Context{})
	if err != nil {
		boolean = false;
	} else {
		err = json.Unmarshal([]byte(resp), &arr)
		if err != nil {
			boolean = false;
		}
		if arr != nil && arr.(map[string]interface{})["returnstatus"] == "Success" {

			mysql.Db.Create(&sms_data)
		}
	}

	defer func(){
		mysql.Db.Close();
	}()
	return boolean,err
}
