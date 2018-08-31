package aes

import (
	"fmt"
	"github.com/adolphlxm/atc/utils/encrypt"
	"testing"
)

func TestAes(t *testing.T) {
	k := []byte("ECU-Upgrade-8369")
	oD := []byte("carlt001")
	crypt, _ := encrypt.AesEncrypt(oD, k)
	fmt.Println(crypt, string(crypt))
}
