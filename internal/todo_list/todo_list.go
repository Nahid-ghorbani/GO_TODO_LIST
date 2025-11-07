package todo_list

import (
	"fmt"

	"github.com/Nahid-ghorbani/GO_TODO_LIST/internal/todo_list/model"
)

func CreateNewTask(args *Args) model.TaskType {
	task := model.TaskType{
		Title:       args.title,
		Description: args.description,
		Status:      model.StatusType(args.status),
	}

	return task
}

func RunTodo() {
	taskValues, err := ParseFlags()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	newTask := CreateNewTask(taskValues)
	fmt.Println("new task", newTask)
}
