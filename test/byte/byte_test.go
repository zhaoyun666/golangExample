package byte

import (
	"encoding/binary"
	hex2 "encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestHex(t *testing.T) {
	hex := fmt.Sprintf("%d", 0x01)
	fmt.Println(hex)
	var a byte
	a = 0x09
	fmt.Println(a & 0xff)
	fmt.Println([]byte{0})
	var b byte
	b = 0x52
	fmt.Println(uint32(b) << 8)
	fmt.Println(20992 >> 8)
	h := []byte{00, 0x01, 0xd6, 0x5a}
	fmt.Println(binary.BigEndian.Uint32(h))

	fmt.Println("V", binary.BigEndian.Uint16([]byte{0x06, 0x7C}))
	fmt.Println("A", binary.BigEndian.Uint16([]byte{0x26, 0xbb}))
}

func TestB(t *testing.T) {
	b := []byte("�����������������")
	fmt.Println("LJ8E3A1M5GE004205", b)
	fmt.Println("LJ8E3A1M5GE004205", string(b))

	fmt.Println([]byte("LJ8E3A1M5GE004205"))

	b = []byte("�����������������")
	fmt.Println("LW4512G44027950P", b)
	fmt.Println(string(b))

	fmt.Println([]byte("LJ8E3A1M5GE004205"))

}

// ecu升级协议数据包
type ProtoPkg struct {
	Header     string `json:"header"`
	Ext        byte   `json:"ext"`
	Platform   byte   `json:"platform"`
	Command    byte   `json:"command"`    // 命令单元，命令标识
	DeviceID   string `json:"deviceid"`   // 设备号
	PayloadLen uint8  `json:"payloadlen"` // 数据报长度
	Payload    []byte `json:"payload"`    // 数据报内容
	End        string `json:"end"`
}

func TestUnmarshal(t *testing.T) {
	s := "71, 69, 84, 32, 47, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 49, 50, 55, 46, 48, 46, 48, 46, 49, 58, 49, 50, 52, 49, 49, 13, 10, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 58, 32, 85, 112, 103, 114, 97, 100, 101, 13, 10, 80, 114, 97, 103"
	sarr := strings.Split(strings.TrimSpace(s), ",")
	b := make([]byte, 0)
	for _, v := range sarr {
		b = append(b, strings.TrimSpace(v)...)
	}
	fmt.Println(b)
	p := &ProtoPkg{}
	err := json.Unmarshal(b, p)
	fmt.Println(err)
	fmt.Println(p)
}

func TestString(t *testing.T) {
	vin := "0xfffffffffffffffff"
	fmt.Println(len(vin) == 17 || strings.Contains(vin, "0xff"))
	buf := make([]byte, 0)
	buf = append(buf, fmt.Sprintf("%10s", "BOSH")...)
	fmt.Println(strings.TrimSpace(string(buf)))
	fmt.Println([]byte("m"), binary.BigEndian.Uint16([]byte{0x00, 0x86}))
	b := make([]byte, 0)
	b = append(b, 0)
	fmt.Println([]byte{0x00, 0x01}, b)
}

// LW4522H37015010E

// LJ8E3A1M4HE008697
// LJ8E3A1M8HE007410

func TestClientIssued(t *testing.T) {
	s := "01"
	p, _ := strconv.Atoi(s)
	fmt.Println(s, p)
	return
	hex := "23 23 c0 fe 4c 57 36 30 36 41 48 33 34 30 30 30 32 30 38 50 30 01 00 22 00 01 10 02 00 00 00 19 00 00 00 17 00 00 00 17 4c 57 36 30 36 41 48 33 34 30 30 30 32 30 38 50 52 ff 8a"
	b := strings.Split(hex, " ")
	for _, v := range b {
		fmt.Println(hex2.DecodeString(v))
	}
}
