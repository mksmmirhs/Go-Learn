package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("enter tax-rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	earning := earningBeforeTax * (1 - taxRate/100)

	fmt.Println("Earning before tax: ", earningBeforeTax)
	fmt.Println("Earning after tax: ", earning)
	fmt.Println("Ratio: ", earningBeforeTax/earning)

}
