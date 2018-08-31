package pointer

import (
	"bytes"
	"encoding/json"
	"github.com/tealeg/xlsx"
	"log"
	"testing"
)

type Response struct {
	Token string
}
type Err struct {
	Code UserErrType
	Msg  string
}
type UserErrType uint32

func TestA1(t *testing.T) {
	x := 1
	p := &x
	log.Println(*p)
	*p = 2
	log.Println("pointer result is :", &p, x)
}

func ErrorMsg(code UserErrType) *Err {
	return &Err{
		Code: code,
		Msg:  Msg[code],
	}
}

var Msg = map[UserErrType]string{
	1: "12",
	2: "23",
}

type version struct {
	Version string
}

func TestVersion(t *testing.T) {
	v := make([]version, 0)
	v1 := version{Version: "1.01.0531"}
	v2 := version{Version: "1.01.0530"}
	v3 := version{Version: "1.01.0529"}
	v = append(v, v1)
	v = append(v, v2)
	v = append(v, v3)

	ver := "0.0.0"
	for _, v := range v {
		log.Println("version compare :", ver, v.Version, bytes.Compare([]byte(ver), []byte(v.Version)))
	}

}

func TestBrandId(t *testing.T) {
	js := `{"2":[],"2359":[],"1":[],"4":[],"1993":[],"2174":[],"7":[],"8":[],"1769":[],"12":[],"14":[],"6":[],"10":[],"15":[],"11":[],"1796":[],"126":[],"9":[],"5":[],"1695":[],"2120":[],"1907":[],"1906":[],"19":[],"17":[],"2362":[],"2579":[],"2382":[],"26":[],"24":[],"2105":[],"2050":[],"1705":[],"23":[],"1706":[],"22":[],"1803":[],"29":[],"32":[],"28":[],"33":[],"2112":[],"30":[],"31":[],"36":[],"35":[],"116":[],"50":[],"1930":[],"123":[],"39":[],"1960":[],"1735":[],"41":[],"40":[],"2135":[],"1768":[],"95":[],"37":[],"2387":[],"38":[],"45":[],"48":[],"51":[],"2403":[],"46":[],"2250":[],"49":[],"2414":[],"2406":[],"1783":[],"1779":[],"54":[],"56":[],"2418":[],"2423":[],"53":[],"55":[],"120":[],"57":[],"118":[],"60":[],"66":[],"67":[],"62":[],"2433":[],"58":[],"64":[],"2542":[],"61":[],"65":[],"63":[],"1774":[],"70":[],"1964":[],"69":[],"72":[],"68":[],"71":[],"2034":[],"2465":[],"121":[],"74":[],"75":[],"1825":[],"124":[],"1834":[],"78":[],"79":[],"82":[],"84":[],"1938":[],"83":[],"94":[],"87":[],"1708":[],"1974":[],"89":[],"90":[],"91":[],"92":[],"93":[],"2478":[],"86":[],"2151":[],"2166":[],"97":[],"100":[],"98":[],"99":[],"2481":[],"101":[],"105":[],"125":[],"102":[],"103":[],"21":[],"108":[],"2185":[],"109":[],"110":[],"107":[],"2507":[],"111":[],"112":[],"115":[]}`
	js = ""
	var brandids map[string][]int64
	json.Unmarshal([]byte(js), &brandids)
	batchids := make(map[string]interface{}, 0)
	inBrandid := false
	for brandid, optionids := range brandids {
		if _, ok := batchids[brandid]; ok && len(optionids) == 0 {
			inBrandid = true
			break
		}
		log.Println(brandid, optionids, batchids, inBrandid)
	}
	log.Println(inBrandid)
}

func TestCreateExcel(t *testing.T) {
	var nav = [6]string{"车架号", "设备号", "总里程", "所属经销商", "下线日期", "活跃状态"}
	// 创建头部
	for _, v := range nav {
		log.Println(v)
	}
	Create()
}

func createExcel() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		log.Println(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "年龄"

	row1 = sheet.AddRow()
	row1.SetHeightCM(1)
	cell = row1.AddCell()
	cell.Value = "狗子"
	cell = row1.AddCell()
	cell.Value = "18"

	row2 = sheet.AddRow()
	row2.SetHeightCM(1)
	cell = row2.AddCell()
	cell.Value = "蛋子"
	cell = row2.AddCell()
	cell.Value = "28"

	err = file.Save("/home/zc/Desktop/Excel/test.xlsx")
	if err != nil {
		log.Println(err.Error())
	}

}
