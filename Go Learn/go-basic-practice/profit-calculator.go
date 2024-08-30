package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate")
	fmt.Scan(&taxRate)

	// earning before tax
	ebt := revenue - expenses

	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit
	fmt.Println("EARNING BEFOR TAX: ",ebt)
	// %v -> any value, %.2f -> 2 width decimal value, %.0f -> 0 decomal place
	fmt.Printf("Profit : %v\n",profit)
	
	fmt.Println(ratio)

}
