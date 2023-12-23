package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {

	var investment float64
	var investmentRate float64
	var years float64

	fmt.Print("Enter your investment:")
	fmt.Scan(&investment)

	fmt.Print("Enter your investment rate:")
	fmt.Scan(&investmentRate)

	fmt.Print("Enter numbers of year:")
	fmt.Scan(&years)

	investmentReturn, futureRealValue := investmentReturn(investment, investmentRate, years)

	fmt.Println("investment return is:", investmentReturn)
	fmt.Println("Real value of investment:", futureRealValue)
}

func investmentReturn(investment, investmentRate, years float64) (ir float64, fr float64) {
	ir = investment * math.Pow((1+investmentRate/100), years)
	fr = ir / math.Pow(1+inflationRate/100, years)
	return ir, fr
}
