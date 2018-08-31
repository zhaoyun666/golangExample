package tests

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"learning-golang-process/test/procWorker/helper"
	"learning-golang-process/test/procWorker/model"
	"learning-golang-process/test/procWorker/utils"
	"os"
	"path"
	"strings"
	"testing"
)

func TestParseGps(t *testing.T) {
	key := "gbzmplant"
	vin := "LJ8E3A1M3GE001349"
	utils.New()
	h := utils.RC.HGet(key, vin)

	res, _ := h.Bytes()
	r := bytes.NewBuffer(res)
	buf6 := make([]byte, 6)
	r.Read(buf6)
	fmt.Println("current Date: ", buf6)
	fmt.Println(fmt.Sprintf("20%02d-%02d-%02d %02d:%02d:%02d", buf6[0], buf6[1], buf6[2], buf6[3], buf6[4], buf6[5]))
	addr := helper.ParseProto(r)

	p := &helper.GpsPosition{
		Position: addr,
	}
	address := helper.ConvertAndResolve(p)
	fmt.Println(address)
}

func TestFetchGPS(t *testing.T) {
	utils.New()
	f, _ := os.Open("/home/work/go/src/learning-golang-process/test/procWorker/tests/vin.txt")
	defer f.Close()
	rd := bufio.NewReader(f)
	res := make([]string, 0)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		res = append(res, line)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestQueryGEO(t *testing.T) {
	utils.New()
	utils.NewGEO()
	res, err := model.GetsGeo()
	if err != nil {
		return
	}

	for _, v := range res {
		a, b, c := fetchGPS(v.Vin, v.Deviceidstring)
		carObdDevice := model.CarObdDevice{
			Vin:     v.Vin,
			Gps:     a,
			Address: b,
			Rectime: c,
		}
		carObdDevice.InsertOne()
	}

}

func task() {
	g, err := utils.RC.GeoPos("275380", "LW4522H37013080M").Result()
	fmt.Println(g, err)

	/*p := &helper.GpsPosition{
		Position: addr,
	}
	address := helper.ConvertAndResolve(p)
	fmt.Println(vin, addr, address, )*/
}

func fetchGPS(vin, dev string) (position, address, rectime string) {
	ge, _ := utils.GEORC.GeoPos("275380", dev).Result()
	var la, lon float64

	for _, v := range ge {
		if v == nil {
			return "", "", ""
		}
		la = v.Latitude
		lon = v.Longitude
	}

	p := &helper.GpsPosition{
		Position: fmt.Sprintf("%.6f,%.6f", la, lon),
	}
	address = helper.ConvertAndResolve(p)

	return p.Position, address, fetchDate(vin)
}

func fetchDate(vin string) string {
	key := "gbzmplant"
	h, _ := utils.RC.HGet(key, vin).Bytes()
	r := bytes.NewBuffer(h)
	buf6 := make([]byte, 6)
	r.Read(buf6)
	return fmt.Sprintf("20%02d-%02d-%02d %02d:%02d:%02d", buf6[0], buf6[1], buf6[2], buf6[3], buf6[4], buf6[5])
}

func TestString(t *testing.T) {
	s := "/data/YeMa-PhotoAlbum/201805/LSA121BL6J20010861527047594067815595.thumbnail.png"
	fmt.Println(path.Base(s))
	t1 := getNewFileName(s, "thumbnail")
	t2 := CheckBaseName(s, "", "")
	fmt.Println(t1, t2)
}

func getNewFileName(source string, tp string) string {
	ext := path.Ext(source)
	fmt.Println(ext, source, tp)
	return strings.Replace(source, fmt.Sprintf(".%s%s", tp, ext), ext, 1)
}

func CheckBaseName(ph string, old, replace string) (name string) {

	baseName := path.Base(ph)
	l := len(baseName)
	if l < 4 {
		return ""
	}

	name = strings.Replace(baseName, old, replace, 1)
	return name
}

func TestDelete(t *testing.T) {
	carObdDevice := model.CarObdDevice{
		Deviceidstring: "LW0011E340000105",
	}
	carObdDevice.Del()
}
