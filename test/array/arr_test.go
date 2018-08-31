package array

import (
	"fmt"
	"testing"
)

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func TestArray(t *testing.T) {
	//q := [3]int{1, 2, 3}
	//q = [4]int{1, 2, 3, 4}
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
}

func TestSlice(t *testing.T) {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2, len(Q2), cap(Q2))
	fmt.Println(summer, len(summer), cap(summer))
	fmt.Println(summer[:5])
}

func TestRev(t *testing.T) {
	v1 := []int{0: 1, 1: 3, 2: 5}
	reverse(v1)
	fmt.Println(v1)

	v2 := []int{1, 2, 3, 4}
	rev(&v2)
	fmt.Println(v2)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rev(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
	}
}
