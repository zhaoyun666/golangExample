package main

import "fmt"

func main() {
    naturals := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 0; ; x++ {
            naturals <- x
        }
    }()

    //Square
    go func() {
        for {
            x := <-naturals
            squares <- x * x
        }
    }()

    //Print
    for {
        fmt.Print(<-squares)
    }

}
