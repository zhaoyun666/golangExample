package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFileEOF(t *testing.T) {
	f, err := os.Open("PCBTest_20180327_PE.abs-1.s19")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
}
