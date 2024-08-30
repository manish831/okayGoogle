package main

import "fmt"

func main() {
	age := 32 //regular variable
	fmt.Println("Age: ", age)
	// pointer topic.
	/*
		1. In one step
		agePointer := 1&age

		2. In 2 step, i.e., declaration & iitialisation
		var agePointer *int
		agePointer = &age

	*/
	agePointer := &age
	fmt.Println("AgePointer: ", agePointer)

	// dereferencing the pointer variable using Asterisk *
	fmt.Println("Age can be retrieved by dereferencing by using *. Age is: ", *agePointer)

	// adultYears := getAdultYears(age) //pass by value
	getAdultYears(agePointer)
	fmt.Println(age)
}

//this function will take int as a arguement.
/*
func getAdultYears(age int) int {
	// create a copy of variable age, and not the same memory.
	return age - 18
}
*/

// this function will take *int pointer as arguement
func getAdultYears(age *int)  {
	*age =  *age - 18
}
