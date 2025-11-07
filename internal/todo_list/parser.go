package todo_list

import (
	"flag"
	"fmt"

	"github.com/Nahid-ghorbani/GO_TODO_LIST/internal/todo_list/model"
)

type Args struct {
	title       string
	description string
	status      string
}

func ParseFlags() (*Args, error) {
	// get flags from command line
	title := flag.String("title", "", "use -title and enter the task title.")
	description := flag.String("description", "", "use -description and enter the task description.")
	status := flag.String("status", string(model.StatusTodo), "use -status and then enter \"todo\" or \"doing\" or \"done\".")

	// parse command and extract values
	flag.Parse()

	// task without title is meaningless
	if *title == "" {
		return nil, fmt.Errorf("missing requires Flag \"title\"")
	}

	// check the status to be in format has defined
	validateStatus := map[string]bool{
		string(model.StatusTodo):  true,
		string(model.StatusDoing): true,
		string(model.StatusDone):  true,
	}

	if !validateStatus[*status] {
		return nil, fmt.Errorf("invalid status: %q (must be todo, doing, or done)", *status)
	}

	return &Args{*title, *description, *status}, nil
}
