package internal

import (
	"errors"
	"time"
)

var (
	// ErrTaskNotFound is an error that is returned when a task is not found
	ErrTaskNotFound = errors.New("task not found")
	// ErrTaskAlreadyExists is an error that is returned when a task already exists
	ErrTaskAlreadyExists = errors.New("task already exists")
)

// Task is a struct that represents a task
type Task struct {
	// ID is the unique identifier of the task
	// - should be unique (auto-incremented)
	ID int
	// Name is the name of the task
	// - should be unique
	Name string
	// Description is the description of the task
	// - 65535 characters max
	Description string
	// Completed is a boolean that represents if the task is completed or not
	// - default value: false
	Completed bool
	// CreatedAt is the date when the task was created
	// - auto-generated
	CreatedAt time.Time
}
