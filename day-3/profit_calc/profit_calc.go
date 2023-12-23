package main

import "fmt"

func main() {

	revenue := getInput("enter revenue: ")
	expenses := getInput("enter expenses: ")
	taxRate := getInput("enter tax rate: ")

	earningBeforeTax, earning := ebtCalculation(revenue, expenses, taxRate)

	fmt.Println("Earning before tax: ", earningBeforeTax)
	fmt.Println("Earning after tax: ", earning)
	fmt.Println("Ratio: ", earningBeforeTax/earning)

}

func getInput(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func ebtCalculation(revenue, expenses, taxRate float64) (ebt, earning float64) {
	ebt = revenue - expenses
	earning = ebt * (1 - taxRate/100)
	return ebt, earning
}
