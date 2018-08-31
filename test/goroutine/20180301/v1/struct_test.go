package v1

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A struct {
	B float64
}
type C struct {
	D string
}

func TestStruct(t *testing.T) {
	A := A{}
	fmt.Println(A)
	C := C{D: "Hello"}
	fmt.Println(C)
	c, _ := json.Marshal(C)
	fmt.Println(c)
}

func TestBinary(t *testing.T) {
	fmt.Println(7 >> 1)
	fmt.Println(11 >> 2)
	fmt.Println(11 >> 3)
}
