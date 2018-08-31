package string

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const (
	a = iota
	b = 2 << iota & 0xff
	c = 2 << iota & 0xff
	d = 2 << iota & 0xff
)

type x8flog struct {
	Endurancemile float64
	Odo           float64
	Avgspeed      float64
	Avgfuel       float64
}

func TestStr(t *testing.T) {
	var status byte = 2
	switch status {
	case 2, 3:
		fmt.Println("xxxxxxxxxxxxxxxxxxxx")
		x8flog := x8flog{
			Endurancemile: 200.3,
			Odo:           108.2,
			Avgspeed:      90.6,
			Avgfuel:       45,
		}
		data := fmt.Sprintf(`enduranceMile=%.1f&ODO=%.1f&avgSpeed=%.1f&avgFuel=%.1f`, x8flog.Endurancemile, x8flog.Odo, x8flog.Avgspeed, x8flog.Avgfuel)
		fmt.Println(data)
		return
	}
	fmt.Println(a, b, c, d)
	enduranceMile, _ := strconv.ParseFloat("", 64)
	fmt.Println(enduranceMile)
}

func TestJoin(t *testing.T) {
	data := make([]string, 0)
	data = append(data, "a=b")
	fmt.Println(strings.Join(data, "&"))
}
