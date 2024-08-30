package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"examples.com/structs/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed!")
		return
	}
	fmt.Println("Saving the note succeeded!")
}
func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// var value string
	// fmt.Scanln(&value) //this can take input single text, let's handle longer user input text

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
