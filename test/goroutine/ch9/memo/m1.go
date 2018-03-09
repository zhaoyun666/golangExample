package main

import (
    "net/http"
    "io/ioutil"
    "time"
    "fmt"
    "log"
)

type Memo struct {
    f Func
    cache map[string]result
}
// Func is the type of the function to momoize.
type Func func(key string) (interface{}, error)

type result struct {
    value interface{}
    err error
}

func New(f Func) *Memo {
    return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
    res, ok := memo.cache[key]
    if !ok {
        res.value, res.err = memo.f(key)
        memo.cache[key] = res
    }
    return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
func incomingURLs() <-chan string {
    ch := make(chan string)
    go func() {
        for _, url := range []string{
            "https://baidu.com",
            "https://godoc.org",
            "https://play.golang.org",
            "http://gopl.io",
            "https://golang.org",
            "https://godoc.org",
            "https://play.golang.org",
            "http://gopl.io",
        } {
            ch <- url
        }
        close(ch)
    }()
    return ch
}

func main() {
    m := New(httpGetBody)
    for url := range incomingURLs() {
        start := time.Now()
        value, err := m.Get(url)
        if err != nil {
            log.Print(err)
        }
        fmt.Printf("%s, %s, %d bytes\n",
            url, time.Since(start), len(value.([]byte)))
    }

}