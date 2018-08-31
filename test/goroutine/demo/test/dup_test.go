package test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestDup1(t *testing.T) {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
