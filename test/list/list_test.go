package ttt

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	L := list.New()
	for i := 0; i < 50; i++ {
		L.PushBack(i)
	}

	for e := L.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
