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

type Todos []TodoList // this is a type slice

// creating . .. . a receiver function that adds title to the todo list
// pointer will be use to access the Todos

func (todos *Todos) add(title string) {
	//code below is where create  what is called * struct literal* something that looks like an object as Go is not object oriented
	todo := TodoList{ // todo is local varriable very common in Go receiver function for effective data manipulation
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
	*todos = append(t[:index], t[index+1:]...) // here the slice was unpacked after the deletion and repacked a common practice in go
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
}
func (todos *Todos) edit(index int, title string) error { //error method was used to get the errors
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌" // this is for a non executed to do
		CompletedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				CompletedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), CompletedAt)
	}
	table.Render()
}
