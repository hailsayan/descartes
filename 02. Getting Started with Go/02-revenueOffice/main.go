package main

import (
	"fmt"
	"math"
)

func main() {
	var income int
	// fmt.Println("enter your annual income: ")
	fmt.Scan(&income)

	tax := 0.0

	if income <= 100 {
		tax = float64(income) * 0.05
	} else if income <= 500 {
		tax = 100*0.05 + float64(income-100)*0.10
	} else if income <= 1000 {
		tax = 100*0.05 + 400*0.10 + float64(income-500)*0.15
	} else {
		tax = 100*0.05 + 400*0.10 + 500*0.15 + float64(income-1000)*0.20
	}

	tax = math.Floor(tax)
	// fmt.Printf("The tax you need to pay is: %d\n", int(tax))
	fmt.Println(int(tax))
}
