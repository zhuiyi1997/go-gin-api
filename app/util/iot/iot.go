package iot

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/zhuiyi1997/go-gin-api/app/config"
	"log"
	"reflect"
)

type Iot struct {
	RegionId string
	AccessKey string
	AccessSecret string
	Client *sdk.Client
}

func NewIot() (Iot,error){
	var poac Iot = Iot{RegionId: config.IotConf["regionId"],AccessKey: config.IotConf["accessKey"],AccessSecret: config.IotConf["accessSecret"]}

	client, err := sdk.NewClientWithAccessKey(poac.RegionId,poac.AccessKey,poac.AccessSecret)

	poac.Client = client

	return poac,err
}

func (iot *Iot) RegisterDevice(device map[string]string) (bool,map[string]string){
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = config.IotConf["iotHost"]
	request.Version = "2018-01-20"
	request.ApiName = "RegisterDevice"
	request.QueryParams["RegionId"] = config.IotConf["regionId"]
	request.QueryParams["ProductKey"] = config.IotConf["productKey"]
	request.QueryParams["IotInstanceId"] = config.IotConf["iotInstanceId"]
	request.QueryParams["DeviceName"] = device["DeviceName"]
	request.QueryParams["Nickname"] = device["Nickname"]

	response, err := iot.Client.ProcessCommonRequest(request)
	m := map[string]string{}
	if err != nil {
		m["error"] = err.Error()
		return false,m;
	}
	resp := response.GetHttpContentString()
	var d interface{}
	err = json.Unmarshal([]byte(resp),&d)
	if err != nil {
		m["error"] = err.Error()
		return false,m;
	}
	log.Println(resp)
	if reflect.ValueOf(d.(map[string]interface{})["Success"]).Bool() {
		for k,v := range d.(map[string]interface{}) {
			if k == "Data" {
				for k1,v1 := range v.(map[string]interface{}) {
					m[k1] = reflect.ValueOf(v1).String()
				}
			}
		}
		return true,m;
	} else {
		m["error"] = reflect.ValueOf(d.(map[string]interface{})["ErrorMessage"]).String()
		return false,m;
	}

}
