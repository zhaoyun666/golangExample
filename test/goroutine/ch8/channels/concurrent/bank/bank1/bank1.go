package main

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)

func Deposit(amount int)  { deposits <- amount }
func Balance() int        { return <-balances }
func WithDraw(amount int) { withdraw <- amount }
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			balance -= amount
		}
	}
}

func main() {
	go teller()
	Deposit(200)
	Deposit(100)
	WithDraw(10)
	WithDraw(50)
	Deposit(1 << 1)
	Deposit(1 << 2)
	fmt.Println(Balance())
}
