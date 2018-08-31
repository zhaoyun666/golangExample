package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var p byte
	var i uint8
	for i = 0; i < 8; i++ {
		p += pc[byte(x>>(i*8))]
	}
	return int(p)
}
func main() {
	fmt.Println(pc)
	/*
	 * 00001000     00011000
	 * 00000111     00010111
	 */
	fmt.Println(24 & 23)
	fmt.Println(PopCount(16))
}
