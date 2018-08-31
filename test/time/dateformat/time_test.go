package dateformat

import (
	"fmt"
	"path"
	"strings"
	"testing"
	"time"
)

const (
	GOTIMELAYOUT     = "20060102150405"
	OUTPUTTIMELAYOUT = "2006-01-02 15:04:05"
	REDISCACHENAME   = "redis"
)

func TestTimeFormat(t *testing.T) {
	fmt.Println(time.Now().Format(GOTIMELAYOUT)[0:8])
}

func TestTimeFormat1(t *testing.T) {
	ns := time.Now().UnixNano() / 1e9
	fmt.Println(time.Unix(ns, 0).Format(GOTIMELAYOUT))
}

func TestNonoFormat(t *testing.T) {
	s := time.Now().UnixNano() / 1e9
	//s = 1523867818215260795 /1e9
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Unix(s, 0).Format(GOTIMELAYOUT))
}

// 1523867818215260795

func TestString(t *testing.T) {
	str := "LSGGF53X1AH2387081523870954825321431.png"
	et := path.Ext(str)
	new_str := strings.Replace(str, et, fmt.Sprintf(".%s%s", "thubnail", et), 1)
	fmt.Println(new_str)
}
