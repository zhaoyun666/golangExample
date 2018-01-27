package string

import (
	"testing"
	"strings"
	"fmt"
	"strconv"
	"math"
)

func TestVersion(t *testing.T) {
	a := "2.30.12"
	b := strings.Split(a, ".")
	ar := make([]int, 0)
	for _, v := range b {
		t1, _ := strconv.Atoi(v)
		ar = append(ar, t1)
	}

	fmt.Println(0<<8 | 2)
	fmt.Println(2<<8 | 30)
	fmt.Println(542<<8 | 12 )
	fmt.Println(ar)
	fmt.Println(14<<8)
	fmt.Println(1<<32)
	fmt.Println(65536>>8)
	//0000111000000000
	fmt.Println(math.Pow(2, 11) + math.Pow(2, 10) + math.Pow(2, 9))
	//fmt.Println(fmt.Sprintf("%d.%d.%d", c>>16&0xff, c>>8&0xff, c&0xff))
}