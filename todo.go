package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// first am creating a stuct to hold the todos
type TodoList struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// let create a slice to hold the struct data. . . . this mimics OOP
//Receiver function plays a lot of roles here

type Todos []TodoList // value of Todos becomes a slice using TodoList structure as a blueprint, now each todo list is stored in a slice indexed

// creating . .. . a receiver function that adds value to the TodoList Struct .... actually the function creates a todolist
// pointer will be use to access the Todos

func (todos *Todos) add(title string) { //A struct literal is how Go lets you create a struct and give its fields values at the same time.
	//code below is where create  what is called * struct literal* something that looks like an object as Go is not object oriented used to asign value to the struct
	todo := TodoList{ // todo is local varriable very common in Go receiver function for effective data manipulation, over here we create a new todo
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	//we will return todos with pointer into a slice using append method
	*todos = append(*todos, todo)
}
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
	}
	return nil

}

// below is function to delete a todo using the index as a method
func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...) // here the slice was unpacked after the deletion and repacked a common practice in go a standard Go pattern to delete things
	return nil
}
func (todos *Todos) toggle(index int) error { //function to toggle a to do
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now() // here uses time function
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted
	return nil
} //edit(parameters) what is inside the bracket after the function name is argument, they are inputs to the fuction, the become argument when the inputs are given to it to execute
func (todos *Todos) edit(index int, title string) error { //error method was used to get the errors
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title

	return nil
}

func (todos *Todos) display() {
	table := table.New(os.Stdout) // os.Stdout → means “print to the terminal”
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌" // this is for a non executed to do
		CompletedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				CompletedAt = t.CompletedAt.Format(time.RFC1123) //RFC1123 → standard human-friendly time format
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), CompletedAt) // strconv.Itoa(index) → converts the integer to a string, because table cells need text
	}
	table.Render()
}
