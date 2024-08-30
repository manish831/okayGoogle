package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)
	fmt.Println(double) // -> address
	fmt.Println(triple) // -> address

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
	// Anonymous Functions
	transformed := transformNumbers(&numbers, func (number int) int {
		return number*2
	})
	
	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}


func createTransformer(factor int) func(int) int{
	// every anonymous function is Closure.
	return func(number int) int {
		return number*factor
	}
}