package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investment float64
	var investmentRate float64
	var years float64

	fmt.Print("Enter your investment:")
	fmt.Scan(&investment)

	fmt.Print("Enter your investment rate:")
	fmt.Scan(&investmentRate)

	fmt.Print("Enter numbers of year:")
	fmt.Scan(&years)

	investmentReturn := investment * math.Pow((1+investmentRate/100), years)
	futureRealValue := investmentReturn / math.Pow(1+inflationRate/100, years)

	fmt.Println("investment return is:", investmentReturn)
	fmt.Println("Real value of investment:", futureRealValue)
}
