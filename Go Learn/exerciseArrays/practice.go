package main

import "fmt"

type Product struct {
	id string
	title string
	price float64
}

func main() {
	// 1. Create a new array, contains threee hobbies.
	var hobbies [3]string = [3]string{"Reading", "Travelling", "Coding"}
	fmt.Println(hobbies)

	/* 2. Output more data about that(1) array
		- The first Element(standalone)
		- The second and third element combined as a new list.
	*/

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	/* 3.   
		Create a slice based on the first 
			element that contains the first and second elements.

		Create that slice in two different ways(i.e. create two slices in )
	*/
	mainHobbies := hobbies[ : 2]
	fmt.Println(mainHobbies)

	// 4. Re-slice the slice from (3) and change it to contain the
	//  second and last element of the original array.

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5. Create a "dynamic array" that contains your coarse goals(at least 2 goals)
			 // dont declare the size --> dynamic
	courseGoals := []string{"Learning Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Backend Developer")
	fmt.Println(courseGoals)

	/*7.
	Bonus: Create a "Product" struct with title, id, price and create a dynamic  list 
		of products(at least 2 products).
	Then add a third product to the existing list of products.
	*/
	products := []Product{
			{
				"first-product",
				"Mine First Product", 
				12.85,
			},
			{
				"Second-product",
				"Mine Second Product", 
				12.85,
			},
		}
	fmt.Println(products)
	newProduct := Product{
		"third-product",
		"A third Product",
		16.0,
	}
	products = append(products,newProduct)
	fmt.Println(products)
}