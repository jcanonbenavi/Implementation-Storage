package internal

type TaskRepository interface {
	// FindByID returns a task by its ID
	FindByID(id int) (task Task, err error)
	// FindAll returns all tasks
	FindAll() (tasks []Task, err error)

	// Save saves a task
	Save(task *Task) (err error)
	// Update updates a task
	Update(task *Task) (err error)
	// Delete deletes a task
	Delete(id int) (err error)
}
