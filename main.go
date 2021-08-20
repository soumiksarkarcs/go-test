package main

import (
	"fmt"
)

func Add(x int, y int) int {
	return x + y
}

func Sub(x int, y int) int {
	return x - y
}

func Mul(x int, y int) int {
	return x * y
}

func Div(x int, y int) int {
	return x / y
}

func main() {
	//fmt.Println("Hello World")
	fmt.Println(Add(2, 3))
	fmt.Println(Sub(12, 3))
	fmt.Println(Mul(2, 3))
	fmt.Println(Div(42, 3))
}
