package main

import (
	"fmt"
	//"math"
)

func main() {
	var num1 float64 = 9
	var num2 float64 = 4
	answer := num1 / num2
	fmt.Printf("%g", answer)
}

func main() {
	var num1 int = 9
	var num2 int = 4
	answer := (num1 % num2) + 5
	fmt.Printf("%d", answer)
}
