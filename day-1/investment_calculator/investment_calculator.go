package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var returnRate = 5.5
	var years = 10
	var futureValue = float64(investmentAmount) * math.Pow((1+returnRate/100), float64(years))
	fmt.Println("future value is:-", futureValue)
}
