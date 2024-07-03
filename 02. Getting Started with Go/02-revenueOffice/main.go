package main

import "fmt"

func main() {
    var incomeInt int
    fmt.Scan(&incomeInt)
    
    var tax float32
    income := float32(incomeInt)
    if income <= 100 {
        tax = income * 0.05
    } else if income > 100 && income <= 500 {
        tax = (100 * 0.05) + ((income - 100) * 0.1)
    } else if income > 500 && income <= 1000 {
        tax = (100 * 0.05) + (400 * 0.1) + ((income - 500) * 0.15)
    } else if income > 1000 {
        tax = (100 * 0.05) + (400 * 0.1) + (500 * 0.15) + ((income - 1000) * 0.2)
    }

    fmt.Println(int(tax))
}