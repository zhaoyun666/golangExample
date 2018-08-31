package db

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	fmt.Println(8501 % 100)
	fmt.Println(8501 / 100)

	fmt.Println(98*87, 99*87, 8601)
	fmt.Println(99*85, 100*85, 8599)
	fmt.Println(99*85, 8500)
	fmt.Println(99*87, 8500)

	a := make([]int, 6)
	fmt.Println(a[4:6])
}

func TestPruduct(t *testing.T) {
	nParseRoutine := 100
	count := 8500
	result := make([]int, 0)
	for i := 1; i <= count; i++ {
		result = append(result, i)
	}

	var n int

	if nParseRoutine > len(result) {
		nParseRoutine = len(result)
		n = 1
	} else {
		n = len(result)/nParseRoutine + 1
	}
	var start int
	for i := 0; i < nParseRoutine; i++ {
		if i*n > count {
			break
		}
		fmt.Println(start, result[i*n:min(int(count), (i+1)*n)])
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}
