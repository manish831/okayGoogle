package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

// for writing on the file, i capitalized the first letter.
type Note struct {
	Title     string    `json : "title"`
	Content   string    `json : "content"`
	CreatedAt time.Time `json : "created_at"`
}

// func (note Note) Display() {
// 	fmt.Printf("Your note title %v had the following content: \n\n%v", note.Title, note.Content)
// 	// lets save into file
// }

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, "  ", "_")
	// to add lowercase
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func New(title, content string) (Note, error) {
	// constructor function contain validation
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input.")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
