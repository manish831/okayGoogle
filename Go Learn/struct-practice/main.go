package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"examples.com/practice/note"
	"examples.com/practice/todo"
)

// saver interface

type saver interface {
	Save() error
}

// creating another interface named displayer..
// type displayer interface {
// 	Display()
// }

// creating embedded interfaces.
type ouputtable interface {
	saver
	// displayer 
	Display() //interface may have function embedded too.
}
func main() {
	printSomething(1)
	printSomething("Print something: Hello")
	printSomething(1.1)
	printSomething('c')

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// todo.Display()
	

	// err = saveData(todo)
	// if err != nil {
	// 	return
	// }
	err = displayAndSave(todo)
	if err != nil {
		return
	}
	// userNote.Display()
	// err = saveData(userNote)
	err = displayAndSave(userNote)
	if err != nil {
		return
	}
	/*
		todo or usernote are first calling Display() 
		function and then saveData() happens...
		So, i am gonna doing to implement another intrerface 
		which first call Display() and then call SaveData().	
	*/

	// Since upon implementing 2 interfaces alg alg, so out of thos one
	// only one can be call -> so i will implement aother interfaces
	// havong these two interfaces embedded.
	// ->Embedded Interfaces.  which is outputtable.

	// err = userNote.Save()
	/*
		- since todo.Save() & userNote.Save() doing same thing, i.e., saving the data
		so why not to make a interface which doing same thing upon sending particular data.
	*/
	// if err != nil {
	// 	fmt.Println("Saving the note failed.")
	// 	return
	// }

	// fmt.Println("Saving the note succeeded!")
}

func printSomething(value interface{}){
	// this function will help me to learn which return anykind of value
	// and take any kind of value as parameter. 

	// <value/var_name> interface{}
	fmt.Println(value)

	/*
	if wanted to print diff diff thing on diff. diff type, use switch
	*/

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println(value)

	// anyother value will be acepted but kuch hoga hi ni, kyuki define hi ni hai.
	}


	// switch case ka alternative niche hai

	intVal, isOk := value.(int)
	if isOk {
		fmt.Println("Integer: ", intVal)
	}

	flaotVal, isOk := value.(float64) 
	if isOk {
		fmt.Println("Float64: ", value)
		// value bhi work kr jayega bcoz had checked ki ye flaot hi hai.
		fmt.Println("Float64: ", flaotVal)
	}

}



func displayAndSave(data ouputtable) error {
	data.Display()
	return saveData(data) 
}
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
