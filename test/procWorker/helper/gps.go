package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	// 地址解析接口
	geoAddress = "http://restapi.amap.com/v3/geocode/regeo?key=%s&location=%s&output=json&extensions=base&batch=false"

	// 坐标转换接口
	gpsConvertAddress = "http://restapi.amap.com/v3/assistant/coordinate/convert?key=%s&locations=%s&coordsys=gps&output=json"
)

const (
	geoKey        = "bed9636f7d663eca555cabfe24e70586"
	gpsConvertKey = "57c9b3d0b6e66ccc8f95885fa2a2446e"
)

type GpsPosition struct {
	Position string // GPS坐标
}

type gpsConvertResponse struct {
	Locations string `json:"locations"`
}

type geoResponse struct {
	Data struct {
		Address string `json:"formatted_address"`
	} `json:"regeocode"`
}

type apiResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	InfoCode string `json:"infocode"`
}

func ConvertAndResolve(p *GpsPosition) string {
	p.Position = sortPos(p.Position)

	// 1. 坐转换
	pos, err := doGpsConvert(p)
	if err != nil {
		fmt.Errorf("[gps] convert(%s) err:%v", p.Position, err)
		return ""
	}

	// 2. 地址解析
	addr, err := doAddressParse(pos)
	if err != nil {
		fmt.Errorf("[gps] parse(%d) err:%v", p.Position, err)
		return ""
	}
	return addr
}

func doHttpGet(url string) ([]byte, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	return buf, err
}

func doGpsConvert(p *GpsPosition) (string, error) {

	urlVal := fmt.Sprintf(gpsConvertAddress, gpsConvertKey, p.Position)
	buf, err := doHttpGet(urlVal)
	if err != nil {
		return "", err
	}

	rsp := &apiResponse{}
	err = json.Unmarshal(buf, &rsp)
	if err != nil {
		return "", err
	}
	if rsp.Status != "1" {
		return "", fmt.Errorf("%s:%s", rsp.InfoCode, rsp.Info)
	}

	ret := &gpsConvertResponse{}
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		return "", err
	}
	return ret.Locations, nil
}

func doAddressParse(pos string) (string, error) {
	urlVal := fmt.Sprintf(geoAddress, geoKey, pos)
	buf, err := doHttpGet(urlVal)
	if err != nil {
		return "", err
	}

	rsp := &apiResponse{}
	err = json.Unmarshal(buf, &rsp)
	if err != nil {
		return "", err
	}
	if rsp.Status != "1" {
		return "", fmt.Errorf("%s:%s", rsp.InfoCode, rsp.Info)
	}

	ret := &geoResponse{}
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		return "", err
	}
	return ret.Data.Address, nil
}

func sortPos(p string) string {
	ps := strings.Split(p, ",")
	if len(ps) < 2 {
		return p
	}
	p1, _ := strconv.ParseFloat(ps[0], 0)
	p2, _ := strconv.ParseFloat(ps[1], 0)
	if p1 < p2 {
		return fmt.Sprintf("%.6f,%.6f", p2, p1)
	}
	return fmt.Sprintf("%.6f,%.6f", p1, p2)
}
