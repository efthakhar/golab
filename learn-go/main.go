package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello dihab")
	// fmt.Println(`hello - dihab`)

	// // working with values
	// var income = 10000
	// var expense = 5000
	// var savings = float64(income) - float64(expense)
	// fmt.Println(savings)
	// fmt.Println(math.Pow(1232, 12))


	number1 := 100
    number1Pointer := &number1
	fmt.Println(number1)
	subtract_ten(number1Pointer)
	fmt.Println(number1)
}

func subtract_ten(num *int) int{
	return *num - 10;
}
