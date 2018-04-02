package main

import (
	//"test/ch7/flag/tempconv"
	"flag"
	"fmt"
)

//var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println("................")
	//fmt.Println(*temp)
}
