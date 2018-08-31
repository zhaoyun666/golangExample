package string

import (
	"testing"
	"github.com/micro/go-log"
	"encoding/hex"
	"io"
	"bytes"
)

func TestArray(t *testing.T) {
	medals := []string{"gold", "silver", "bronze"}
	m1 := make(map[string][]string, 0)
	item := make([]string, 0)
	item = append(item, "Jack", "Tom")
	m1["gold"] = item
	log.Log(medals[0], m1)
}

func TestIntArray(t *testing.T) {
	// int type
	var a []int
	var b [3]int
	log.Logf("a array length:%d, b array length:%d", len(a), len(b))

	// init int array
	var q  = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	log.Log("q array :", q, len(q), cap(q))
	r := [...]int{99:-1}
	log.Log("init int array", r)
	log.Log("bytes ->", []byte{0, 2})
	var b0 byte = 0
	iTa := 1
	iTire := 1
	b0 = b0 | byte(iTa & 0x1)
	log.Log("one step: ", b0)
	b0 = b0 | byte(iTire << 1 & 0x2)
	log.Log("two step: ", 1<<1, b0)

	//s := []byte{86 66 97 116 61 49 50 46 51 38 67 83 81 61 51 49 38 118 101 114 115 105 111 110 61 86 51 46 48 50 46 48 56 48 55 38 86 77 61 49}
	//log.Log([]byte(s))
}

func TestBytes(t *testing.T) {

	x := "232305fe2a2a7777772e6361726c742e636f6d2a2a01002912080e110e300001667a7a686f6e677461690000667a7a686f6e67746169000000000000000000000196"
	c, _ := hex.DecodeString(x)
	var rBuf [32]byte
	r := bytes.NewReader(c)

	n, err := io.ReadFull(r, rBuf[:24])
	log.Log(c)
	log.Log(rBuf)
	log.Log(n, err)
}

// 189901304579

// LW1332I30004820U
// LW605BI09000010X
// 898602B1111530049428
// LW605BI18008690X

/**
LW2332I150003101  2021-7-10 17:04　已完成
LW2332I15000330H  2021-7-28 18:50　已完成
LW2332I15000340P  2021-6-17 13:05　已完成
LW2332I15000350X  2021-8-3 13:49　已完成
LW2332I150003605  2021-7-22 12:12　已完成
LW2332I15000390T  2021-7-31 15:46　已完成
 */

 // 18753695842

 /**
 需要恢复
 设备号
 LW605BI010010621
 原始可用车架号和ｐｉｎ码
 LJ8F3D5H8JE027584 7DCDFDA5
 替换新车架号和ｐｉｎ码
 LJ8F3D5H0JE027241 84B67E55
  */
  /**
  需要恢复
  暂时移除的数据
  设备号
  LW5512H07066300O
  车架号和ｐｉｎ码
  LJ8F3D5H0JE027241 84B67E55
   */