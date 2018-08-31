package excel

import (
	"encoding/csv"
	"github.com/micro/go-log"
	"os"
	"strings"
	"sync"
	"time"
	"fmt"
)

var wg sync.WaitGroup

func Csv(obd []*ObdDevice) {
	result := obd

	var n int
	nParseRoutine := 100
	count := 1000
	if nParseRoutine > len(result) {
		nParseRoutine = len(result)
		n = 1
	} else {
		n = len(result)/nParseRoutine
	}

	length := len(result)
	for i := 0; i < nParseRoutine; i++ {
		if i*n > length {
			break
		}
		GoRoutine(wg, result[i*n:min(int(count), (i+1)*n)], i)
	}
	wg.Wait()
	time.Sleep(time.Second*6)
}

func GoRoutine(wg sync.WaitGroup, obd []*ObdDevice, n int) {
	wg.Add(1)
	go func() {
		f, err := os.Create(fmt.Sprintf("/home/zc/Desktop/Excel/%d.csv", n+1))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		f.WriteString("\xEF\xBB\xBF")

		header := []string{"车架号", "设备号", "CCID", "pin码"}
		w := csv.NewWriter(f)
		data := [][]string{
			header,
		}
		for _, v := range obd {
			item := []string{}
			item = append(item, v.Vin, v.Deviceidstring, v.Ccid, strings.TrimSpace(v.PinNumber))
			data = append(data, item)
		}
		w.WriteAll(data)
		w.Flush()
		wg.Done()
	}()
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}