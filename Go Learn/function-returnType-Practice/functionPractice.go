package main

import "fmt"

func details() (string, string) {
	name := "Manish"
	var fatherName string = "PitaJi"
	return name, fatherName
}
// since both name and fatherName is of type string ==> so have mentioned once.
func doOutput(name, fatherName string){
	fmt.Printf("My name is: %v and My father is my %s \n", name, fatherName)
}
func myDetails(name string, age int){
	fmt.Printf("Name is %v, and age is %d\n", name, age);
}

func main() {
	name, fatherName := details()
	var myage int = 23

	doOutput(name, fatherName);

	fmt.Print(name,"\n", fatherName, "\n")
	myDetails(name, myage);
}