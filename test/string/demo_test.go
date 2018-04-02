package string

import (
	"encoding/json"
	"fmt"
	"github.com/adolphlxm/atc/logs"
	"strconv"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	str := "1101070f1c391001002402002403001e040d057777772e72656469732e636f6d0618eb0756312e30000856322e3039093c0a001e0b001e0c0d0d100e7777772e78787878787878782e636f6d0f69891001"

	fmt.Println(fmt.Sprintf("%d", 0x1c))
	fmt.Println(fmt.Sprintf("%d", 0x82))
	fmt.Println(fmt.Sprintf("%d", 0xA08))
	fmt.Println(fmt.Sprintf("%d", 014))
	fmt.Println(intToByteArray(2568))
	fmt.Println(fmt.Sprintf("%c", 0x67))
	fmt.Println(fmt.Sprintf("0X%02X", 130))
	fmt.Println(50 / 3)

	newReport := make([]string, 0)
	tb := `"eyJjb2RlIjoyODAxLCJyZXF1ZXN0IjoiMTAwL3JlbW90ZS9tb2RpZnlDYXIiLCJtc2ciOiJcdThiZTVcdThiYmVcdTU5MDdcdTRlMGRcdTY1MmZcdTYzMDFcdTYwYThcdTc2ODRcdTU0YzFcdTcyNGMiLCJkYXRhIjp7fX0="`
	ReportDesc := make([]byte, 0)
	report_des_json_string := fmt.Sprintf(`%s`, tb)
	err := json.Unmarshal([]byte(report_des_json_string), &ReportDesc)
	fmt.Println(err)
	str = string(ReportDesc)

	if strings.Contains(str, "=") {
		tp := strings.Split(str, "&")
		for _, v := range tp {
			temp := strings.Split(v, "=")
			newReport = append(newReport, temp[0]+":"+temp[1])
		}
	} else {
		str = strings.TrimPrefix(str, `{"`)
		str = strings.TrimPrefix(str, `}"`)
		tp := strings.Split(str, ",")
		for _, v := range tp {
			if strings.Contains(v, "\\u") {
				sp := strings.Split(v, ":")
				v = "msg:" + formatUnicode(sp[1])
			}
			newReport = append(newReport, v)
		}
	}

}

/**
 * 将unicode码装换为汉字
 */
func formatUnicode(str string) string {
	str = str[1 : len(str)-1]
	sUnicodev := strings.Split(str, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		fmt.Println(err)
		if err != nil {
			logs.Error("Convert Unicode int32 to int64 Error:" + err.Error())
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}

func intToByteArray(a int) []byte {
	ret := make([]byte, 4)
	ret[3] = byte(a & 0xFF)
	ret[2] = byte((a >> 8) & 0xFF)
	ret[1] = byte((a >> 16) & 0xFF)
	ret[0] = byte((a >> 24) & 0xFF)
	return ret
}
