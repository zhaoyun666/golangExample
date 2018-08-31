package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Println(f)
	c := w.(*bytes.Buffer)
	fmt.Println(c)
}
