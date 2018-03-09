package v1

import (
    "testing"
    "fmt"
    "encoding/json"
)

type A struct {
    B float64
}
type C struct {
    D string
}
func TestStruct(t *testing.T){
    A := A{}
    fmt.Println(A)
    C := C{D:"Hello"}
    fmt.Println(C)
    c, _ := json.Marshal(C)
    fmt.Println(c)
}