package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5
	// taking input from console.
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// taking input year
	fmt.Print("Years: ")
	fmt.Scan(&years)

	// expected returnRate
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue /math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}