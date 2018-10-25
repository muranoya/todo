package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"time"
)

type Todo struct {
	Items struct {
		Id      int       `json:"id"`
		Msg     string    `json:"msg"`
		Detail  string    `json:"detail"`
		Created time.Time `json:"created"`
	} `json:"items"`
}

func getJsonDataPath() string {
	if u, err := user.Current(); err != nil {
		panic(err)
	} else {
		return u.HomeDir + "/.todo.json"
	}
}

func loadTodoData(path string) (Todo, error) {
	var todo Todo
	if bytes, err := ioutil.ReadFile(path); err != nil {
		return todo, err
	} else {
		if err := json.Unmarshal(bytes, &todo); err != nil {
			return todo, err
		} else {
			return todo, nil
		}
	}
}

func saveTodoData(path string, todo Todo) (Todo, error) {
	var todo Todo
	return todo, fmt.Errorf("not implemented.")
}

func addTodo() {
	loadTodoData(getJsonDataPath())
}

func showTodo() {
	loadTodoData(getJsonDataPath())
}

func doneTodo() {
	loadTodoData(getJsonDataPath())
}

func main() {
	add := flag.NewFlagSet("add", flag.ExitOnError)
	add.String("detail", "", "Detail of todo.")

	show := flag.NewFlagSet("show", flag.ExitOnError)

	done := flag.NewFlagSet("done", flag.ExitOnError)
	done.Bool("all", false, "Purge all todos.")

	help := flag.NewFlagSet("help", flag.ExitOnError)

	printUsage := func() {
		add.Usage()
		show.Usage()
		done.Usage()
		help.Usage()
	}

	if len(os.Args) <= 1 {
		fmt.Println("Needs sub-command.")
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		add.Parse(os.Args[2:])
		addTodo()
	case "show":
		show.Parse(os.Args[2:])
		showTodo()
	case "done":
		done.Parse(os.Args[2:])
		doneTodo()
	case "help":
		help.Parse(os.Args[2:])
		printUsage()
	default:
		fmt.Printf("%q is not valid sub-command.\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}
