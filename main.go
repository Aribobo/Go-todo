package main

func main() {
	todos := Todos{} //creates  empty variable for todo list
	storage := NewStorsge[Todos]("todo.json")
	storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	// todos.add("buy milk")  //adds a todo to the empty todo list
	// todos.add("buy bread") //adds another todo list
	// fmt.Printf("%+v\n\n", todos) //to print out the to do information
	// todos.delete(0)              //deletes a todo
	// fmt.Printf("%+v", todos)     //prints the remaining todo
	// todos.toggle(0)
	// todos.toggle(1)
	// todos.print()
	storage.Save(todos)
}
