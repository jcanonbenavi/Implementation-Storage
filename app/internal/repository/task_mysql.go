package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/jcanonbenavi/app/internal"
)

// TaskMySQL is a struct that represents a MySQL task repository
type TaskMySQL struct {
	// db is the database connection
	db *sql.DB
}

// FindByID returns a task by its ID
func (t *TaskMySQL) FindByID(id int) (task internal.Task, err error) {
	// query to find a task by its ID
	query := "SELECT `id`, `name`, `description`, `completed`, `created_at` FROM `tasks` WHERE `id` = ?"
	// execute the query
	row := t.db.QueryRow(query, id)
	if row.Err() != nil {
		err = row.Err()
		return
	}
	// scan the result into the task struct
	err = row.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt)
	if err != nil {
		// check if the error is sql.ErrNoRows
		if err == sql.ErrNoRows {
			// if it is, return internal.ErrTaskNotFound
			err = internal.ErrTaskNotFound
			return
		}
		return
	}

	return

}

// Save saves a task
func (t *TaskMySQL) Save(task *internal.Task) (err error) {
	// query to insert a task
	query := "INSERT INTO `tasks` (`name`, `description`, `completed`, `created_at`) VALUES (?, ?, ?, ?)"
	// execute the query
	//exec is a method of the sql.DB struct that executes a query without returning any rows
	// it receives a query and the arguments for the query as parameters
	result, err := t.db.Exec(query, task.Name, task.Description, task.Completed, task.CreatedAt)
	if err != nil {
		// check if the error is a mysql error
		var mysqlErr *mysql.MySQLError
		//errors.As is a function that checks if an error is of a specific type
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrTaskAlreadyExists
				return
			default:
				return
			}
		}
		return
		// mysqlErr, ok := err.(*mysql.MySQLError)
		// if ok {
		// 	if mysqlErr.Number == 1062 {
		// 		err = internal.ErrTaskAlreadyExists
		// 		return
		// 	}
		// }
		// return
	}
	// get the last inserted ID
	//LastInsertId is a method of the sql.Result struct that returns the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return
	}
	// set the ID of the task
	task.ID = int(lastID)
	return
}

// Update updates a task
func (t *TaskMySQL) Update(task *internal.Task) (err error) {
	// query to update a task
	_, err = t.db.Exec(
		"UPDATE `tasks` SET `name` = ?, `description` = ?, `completed` = ? WHERE `id` = ?",
		task.Name, task.Description, task.Completed, task.ID,
	)
	if err != nil {
		// check if the error is a mysql error
		var mysqlErr *mysql.MySQLError
		//errors.As is a function that checks if an error is of a specific type
		if errors.As(err, &mysqlErr) {
			// check if the error is a duplicate entry error
			switch mysqlErr.Number {
			// if it is, return internal.ErrTaskAlreadyExists
			case 1062:
				err = internal.ErrTaskAlreadyExists
				return
			default:
				return
			}
		}
		return
	}

	return
}

// Delete deletes a task
func (t *TaskMySQL) Delete(id int) (err error) {
	// // init transaction
	// tx, err := t.db.Begin()
	// defer func() {
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return
	// 	}
	// 	err = tx.Commit()
	// }()

	// execute the query
	result, err := t.db.Exec(
		"DELETE FROM `tasks` WHERE `id` = ?",
		id,
	)
	if err != nil {
		return
	}

	// get rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	// check rows affected
	if rowsAffected != 1 {
		switch {
		case rowsAffected == 0:
			err = internal.ErrTaskNotFound
		default:
			err = errors.New("more than one row affected")
		}
		return
	}
	// if rowsAffected == 0 {
	// 	err = internal.ErrTaskNotFound
	// 	return
	// } else if rowsAffected > 1 {
	// 	err = errors.New("more than one row affected")
	// 	return
	// }

	return
}

// FindAll returns all tasks
func (t *TaskMySQL) FindAll() (tasks []internal.Task, err error) {
	query := "SELECT `t.id`, `t.name`, `t.description`, `t.completed`, `t.created_at` FROM `tasks t`"
	rows, err := t.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	// iterate over the rows
	for rows.Next() {
		// scan the row into a task struct
		var task internal.Task
		//Scan is a method of the sql.Rows struct that scans the current row into the given variables
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return
		}
		// append the task to the tasks slice
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
