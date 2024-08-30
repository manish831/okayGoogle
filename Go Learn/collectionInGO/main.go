package main

import "fmt"

// allias in map
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// userName := []string{}
	// fmt.Println(cap(userName))

	userName := make([]string, 2, 5) // 1st -> type, 2nd -> intital size, 3rd -> max capacity(optional)
	// phle hi 2 khali element aa gya hai.

	// isko htane ke liy
	userName[0] = "Openmynz"
	userName = append(userName, "Max")
	userName = append(userName, "Manish")
	fmt.Println(userName)
	fmt.Println(cap(userName))

	// courseRating := map[string]float64{}
	/*
		while defining map with above method, go usually do asign multiple
		times memory.
		so, bar bar memory allocate krna ke bachane ke liy, make() function
		ka use krte hai!!
	*/

	// courseRating := make(map[string]float64, 3)
	// let's use allias
	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.7
	courseRating["React"] = 4.8
	courseRating["angular"] = 4.71
	courseRating["node"] = 4.65 // yaha pe memory reallocate hoyga, becoz
	//previously we had made map of memory size 3.
	courseRating.output()
	fmt.Println(courseRating)

	// for loops
	for index, value := range userName {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}
	// for maps
	for key, value := range courseRating {
		fmt.Println("Key", key)
		fmt.Println("Value", value)
	}

}
