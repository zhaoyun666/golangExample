package string

import (
    "testing"
    "fmt"
)

const (
    a = iota
    b = 2 << iota & 0xff
    c = 2 << iota & 0xff
    d = 2 << iota & 0xff
)
func TestStr(t *testing.T) {
    fmt.Println(a, b, c, d)
}
