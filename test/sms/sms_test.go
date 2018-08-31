package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSendSms(t *testing.T) {
	message := "【车乐盒子】您的验证码是:5281"
	q := url.QueryEscape(message)
	fmt.Println(q)
}
