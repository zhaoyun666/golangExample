package work

import (
	"sync"
	"testing"
	"time"
	"log"
	"os"
	"bufio"
	"io"
)

var (
	wg sync.WaitGroup
	out = make(chan string, 0)
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 3; i++ {
		worker()
	}
	wg.Wait()
}

func worker() {
	wg.Add(1)
	go func() {
		log.Print(time.Now().Unix())
		wg.Done()
	}()
}

func TestRead(t *testing.T) {
	go readFile()
	go outString()
	wg.Wait()
	time.Sleep(3*time.Second)
}

func outString() {
	wg.Add(1)
	for v := range out {
		log.Print(v)
	}
	wg.Done()
}

func readFile() {
	wg.Add(1)
	f, _ := os.Open("./out.txt")
	defer f.Close()
	r := bufio.NewReader(f)
	for{
		if line, _, err := r.ReadLine(); err != io.EOF {
			out <- string(line)
		}else{
			return
		}
	}
	wg.Done()
}
