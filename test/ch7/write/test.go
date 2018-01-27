package main

import (
    "io"
    "os"
    "bytes"
    "fmt"
)

func main() {
    var w io.Writer
    w = os.Stdout
    if w != nil {
        w.Write([]byte("Hello"))
    }

    var w1 = new(bytes.Buffer)
    w1.Write([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
    fmt.Println(w1.Cap(), w1.Len())
    fmt.Sprintf("invalid type %v", "xx")
}
