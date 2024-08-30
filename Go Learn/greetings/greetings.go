package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person, or retrurn errors.

func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	// message := fmt.Sprintf(randomFormat(), name)
	
	// creates a message suing a random format.
	message := fmt.Sprintf(randomFormat())
	return message, nil
}
// Hellos returns a map that associates each of thr named pepole with a greeting message.

func Hellos (name []string) (map[string]string, error){
	// A map to associates name with messages\
	messages := make(map[string]string)
	/*
		initialize a map with the following syntax: 
		make(map[key-type]value-type)
	*/

	// Loop through the received slice of names, calling the Hello fucntion to get a messages for each name.

	for _, name := range name { // range return 2 values: the index of current item in the loop and a copy of the item's value
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associates the retrieved messag with the name.
		messages[name] = message
	}
	return messages, nil
}


// randomFormat return one of a set of greeting messages. The returned
// messages is selected at random.

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// return a randomly selected message format by specifyig
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
