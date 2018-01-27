package string

import (
	"testing"
	"fmt"
	"bytes"
)

func TestByte(t *testing.T) {
	str := "1101070f1c391001002402002403001e040d057777772e72656469732e636f6d0618eb0756312e30000856322e3039093c0a001e0b001e0c0d0d100e7777772e78787878787878782e636f6d0f69891001"
	sByte := make([]byte, 0)
	for i := 0; i < len(str)/2; i++ {
		sByte = append(sByte, str[i], str[i+1])
	}
	fmt.Println(fmt.Sprintf("%d", 0x11))
	bf := bytes.Buffer{}
	b, _ := bf.WriteString(str)
	fmt.Println(b)
	fmt.Println(bf.ReadByte())
	fmt.Println(24<<8 | 254)
	fmt.Println(105<<8 | 137)
	fmt.Println(401 | 0xff)
}
