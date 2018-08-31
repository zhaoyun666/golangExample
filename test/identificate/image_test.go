package identificate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/micro/go-log"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

// 人脸对比测试用例
func TestCompareFace(t *testing.T) {
	url := "http://recognition.image.myqcloud.com/face/compare"
	token := Sign()
	data := Request{
		AppId: appid,
		UrlA:  "https://www-linewin-cc.oss-cn-hangzhou.aliyuncs.com/Test/1/timg.jpg",
		UrlB:  "https://www-linewin-cc.oss-cn-hangzhou.aliyuncs.com/Test/1/x.jpg",
	}
	DoRequest(url, token, data)
}

// 人脸对比测试用例
func TestDetectFace(t *testing.T) {
	url := "http://recognition.image.myqcloud.com/face/livedetectpicture"
	token := Sign()
	data := Request{
		AppId: appid,
		Url:   "https://www-linewin-cc.oss-cn-hangzhou.aliyuncs.com/Test/1/1x.jpg",
	}
	DoRequest(url, token, data)
}

var (
	appid      = "10142641"
	bucket     = ""
	secret_id  = "AKIDvujNq7JwAD9ye9ObZTROHAxrmZHyOI1l"
	secret_key = "pcGwqot3JbxCJG3g9Xr48DGCzjKn7ILI"
	expired    = time.Now().Unix() + 86400
	current    = time.Now().Unix()
)

type Request struct {
	AppId string `json:"appid"`
	Url   string `json:"url"`
	UrlA  string `json:"urlA"`
	UrlB  string `json:"urlB"`
}

func auth() string {

	return ""
}

func Sign() string {
	rand.Seed(time.Now().UnixNano())
	rdm := rand.Uint32()
	srcStr := fmt.Sprintf("a=%s&b=%s&k=%s&e=%d&t=%d&r=%d&f=", appid,
		bucket, secret_id, expired, current, rdm)
	sha1Str := fmt.Sprintf("%s%s", Hma(srcStr, secret_key), srcStr)
	signStr := base64.StdEncoding.EncodeToString([]byte(sha1Str))
	log.Log(signStr, len(signStr))
	return string(signStr)
}

func Hma(data, key string) string {
	k := []byte(key)
	d := []byte(data)
	mac := hmac.New(sha1.New, k)
	mac.Write(d)
	return string(mac.Sum(nil))
}

func DoRequest(url, token string, data Request) {

	postData, _ := json.Marshal(data)

	reader := bytes.NewReader(postData)
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, reader)
	req.Header.Add("Content-Type", "application/json")
	req.ContentLength = int64(len(postData))
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Errorf("response err: %v", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logf("read body err: %v", err)
		return
	}
	log.Log(string(b))
}

// 15888911272
