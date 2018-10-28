package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/muranoya/todo/todo"
)

func getJSONDataPath() string {
	if u, err := user.Current(); err != nil {
		panic(err)
	} else {
		return u.HomeDir + "/.todo.json"
	}
}

func loadTodoData(path string) (todo.Todo, error) {
	var todo todo.Todo
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return todo, err
	}

	if err := json.Unmarshal(bytes, &todo); err != nil {
		return todo, err
	}
	return todo, nil
}

func saveTodoData(path string, todo todo.Todo) error {
	bytes, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func addTodo(msg, detail string) error {
	// Ignore error because addTodo is use to create new JSON file
	// when the JSON file has not exist yet.
	todo, _ := loadTodoData(getJSONDataPath())
	if err := todo.AddTodo(msg, detail); err != nil {
		return err
	}
	if err := saveTodoData(getJSONDataPath(), todo); err != nil {
		return err
	}
	return nil
}

func showTodo(allshow bool) error {
	todo, err := loadTodoData(getJSONDataPath())
	if err != nil {
		return err
	}

	todo.PrintItems(allshow)
	return nil
}

func doneTodo(alldone bool, doneid int, uncmplid int, clean bool) error {
	todo, err := loadTodoData(getJSONDataPath())
	if err != nil {
		return err
	}

	if alldone {
		todo.AllDone()
	}

	if doneid >= 0 {
		if err := todo.SetStatus(uint(doneid), true); err != nil {
			return err
		}
	}

	if uncmplid >= 0 {
		if err := todo.SetStatus(uint(uncmplid), false); err != nil {
			return err
		}
	}

	if clean {
		todo.Clean()
	}

	if err := saveTodoData(getJSONDataPath(), todo); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Needs sub-command, add, show, done.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		var msg string
		var detail string
		add := flag.NewFlagSet("add", flag.ExitOnError)
		add.StringVar(&msg, "msg", "", "Message of todo.")
		add.StringVar(&detail, "detail", "", "Detail of todo.")
		add.Parse(os.Args[2:])
		if err := addTodo(msg, detail); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "show":
		var all bool
		show := flag.NewFlagSet("show", flag.ExitOnError)
		show.BoolVar(&all, "all", false, "Show all todos.")
		show.Parse(os.Args[2:])
		if err := showTodo(all); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "done":
		var all bool
		var doneid int
		var uncmplid int
		var clean bool
		done := flag.NewFlagSet("done", flag.ExitOnError)
		done.BoolVar(&all, "all", false, "All incomplete todos makes done.")
		done.IntVar(&doneid, "id", -1, "Set done.")
		done.IntVar(&uncmplid, "unset", -1, "Make finished todo incomplete.")
		done.BoolVar(&clean, "clean", false, "Clean todos.")
		done.Parse(os.Args[2:])
		if err := doneTodo(all, doneid, uncmplid, clean); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Printf("%q is not valid sub-command.\n", os.Args[1])
		os.Exit(1)
	}
}
