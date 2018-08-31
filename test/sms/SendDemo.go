package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	GOTIMELAYOUT = "20060102150405"
	base64Table  = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func main() {
	//APIKey(用户名)
	apiKey := "501112870001"
	//APISecret(密码)
	apiSecret := "FB8D49598725FBFF17DFE4E3E2F6289D1"
	//接受短信的手机号
	mobile := "13720655738"
	//短信内容(【签名】+短信内容)(编码utf8)，系统提供的测试签名和内容，如需要发送自己的短信内容请在启瑞云平台申请签名和模板。
	message := "【车乐盒子】您的验证码是:5381"
	pw := makeSign(apiKey, apiSecret, message)

	smsUrl := fmt.Sprintf("http://api.qirui.com:7891/mt?dc=15&un=%s&pw=%s&da=%s&sm=%s&tf=3&rf=2&rd=1", apiKey, apiSecret, mobile, url.QueryEscape(message))
	fmt.Println(smsUrl, pw)
	sendMessage(smsUrl)
}

func makeSign(apikey, apiSecret, message string) string {
	un := []byte(apikey)
	pw := []byte(apiSecret)
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, uint64(time.Now().Unix()))
	sm := []byte(message)
	c := md5.New()
	c.Write(un)
	c.Write(pw)
	c.Write(ts)
	c.Write(sm)
	e := c.Sum(nil)
	return coder.EncodeToString(e)
}

func sendMessage(smsUrl string) {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", smsUrl, nil)
	reqest.Header.Set("Accept", "application/json, text/plain, */*")
	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		//打印返回结果
		fmt.Println(bodystr)
	}
}
