package model

type StatusType string

const (
	StatusTodo  StatusType = "todo"
	StatusDoing StatusType = "doing"
	StatusDone  StatusType = "done"
)

type TaskType struct {
	Title       string
	Description string
	Status      StatusType
}
