package main

import (
	"fmt"
)

func main() {
	var salary float64
	var income float64

	fmt.Print("Your salary: ")
	fmt.Scan(&salary)

	fmt.Print("Income (Deposited in bank): ")
	fmt.Scan(&income)

	taxRate := 100 - (income / salary) * 100

  formatedTaxRate := fmt.Sprintf("Taxes paid: %.0f%%\n", taxRate)

	fmt.Print(formatedTaxRate)
}
