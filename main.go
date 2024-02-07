package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gnana997/todo-cli-go/todo"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new task")
	complete := flag.Int("complete", -1, "mark a task as done")
	del := flag.Int("del", -1, "delete a task")
	list := flag.Bool("list", false, "list all tasks")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Test Task")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Todo List: %+v\n", todos)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		err := todos.Remove(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Todo List: %+v\n", todos)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *list:
		todos.List()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
