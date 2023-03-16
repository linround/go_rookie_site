package main

import (
	"fmt"
	"time"
)

// Package bank implements a bank with only one account.
var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

func main() {
	// Alice:
	go func() {
		fmt.Println("A1")
		fmt.Println("A2")
	}()

	// Bob:
	go func() {
		fmt.Println("B1")
	}() // B
	time.Sleep(2 * time.Second)
}
