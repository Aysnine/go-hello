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
}
