package todo_list

import (
	"encoding/json"
	"fmt"
	"os"

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

func WriteOnTheJsonFile(task model.TaskType, fileName string) error {
	var tasks []model.TaskType

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to open the file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err.Error() != "EOF" {
		return fmt.Errorf("failed to decode tasks from file: %v", err)
	}

	tasks = append(tasks, task)
	file.Seek(0,0)
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&tasks); err != nil {
		return fmt.Errorf("failed to encode tasks to file: %v", err)
	}

	return nil
}

func RunTodo() {
	taskValues, err := ParseFlags()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	newTask := CreateNewTask(taskValues)
	
	err = WriteOnTheJsonFile(newTask, "tasks.json")
	if err != nil {
		fmt.Println("error writing task on json file:", err)
		return
	}

	fmt.Println("new task", newTask)
}
