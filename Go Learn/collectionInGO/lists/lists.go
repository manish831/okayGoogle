package list

import "fmt"

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

func main() {
	var productNames [4]string = [4]string{"OpenMynz"}
	prices := []float64{10.99, 9.99}
	// fmt.Println(prices);

	prices[1] = 9.99
	prices = append(prices, 5.99, 12.9, 29.99, 100.10)

	prices = prices[1: ]
	fmt.Println(prices);

	discountedPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountedPrices...) // here ... (three dots)are used for unpacking the variable.

	fmt.Println(prices)
	// slices
	featuredPrices := prices[1 : 3]
	fmt.Println(featuredPrices)
	// another method of slicing is 
	/*
		1. [ : 4] ==> from starting to index 4(excluding)
		2. [2 : ] ==> from 2nd index(including) to end of array.
		Note 1.: It doesn't support negative as of Python.
		Note 2.: while creting slice, we don't creating a new copy array.
		But we do reference some partof the original array..
		 --> Memory Efficient.		

		len() -> no. of element currently into the Data Structure.\
		cap() -> capacity of that data structure.
	*/
	highlightedPrice := featuredPrices[:1]
	fmt.Println(highlightedPrice)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrice), cap(highlightedPrice))
	productNames[2] = "A Softlabs"
	fmt.Println(productNames)
	fmt.Println(prices[2])
}