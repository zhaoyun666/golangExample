package test

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once

func TestSyncOnce(t *testing.T) {
	doPrint()
	doPrint()
}

func setup() {
	fmt.Println("Cougratulations for you")
}

func doPrint() {
	once.Do(setup)
}
