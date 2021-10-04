package main

import (
	"fmt"

	"github.com/Aysnine/go-hello/reusable"
)

func main() {
	reusable.Calc()
	reusable.Calc2()
	fmt.Println("hello golang")

	var greeting string = "hello string"
	fmt.Println(greeting)

	greeting2 := "hello var"
	fmt.Println(greeting2)

	greeting2 = "hello var other value"
	fmt.Println(greeting2)

	var num1 int
	fmt.Println("hello number", num1)
}
