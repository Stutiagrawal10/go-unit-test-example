package main

import (
	"fmt"
	"new_project/big_hello"
	"new_project/random"
)

func main() {
	bigHello := bighello.New(random.New())
	fmt.Println(bigHello.DoSomething(1, 2))
}