package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predifened logger, including the log entry prefix
	// and a flag to disable printing the time, source file, and a line number

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Manish", "Ambika", "Sanskriti", "Sophiya"}
	// request a greeting messages for the names.
	messages, err := greetings.Hellos(names)

	// Request a greeting message
	// messages, err := greetings.Hello(names)
	// message, err := greetings.Hello("") // we pass empty string to check any errors.
	// If any error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}
	// if no error was returned, print the returned message to the console
	// Get a greeting message and print it.
	// message := greetings.Hello("Gladys")
	fmt.Println(messages)
}
