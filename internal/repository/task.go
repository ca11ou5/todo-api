package repository

import (
	"fmt"
	"todo-list/internal/entity"
)

func (r *Repository) InsertTask(task *entity.Task) (int64, error) {
	var id int64
	err := r.QueryRow("INSERT INTO tasks(title, description, date, completed) VALUES ($1, $2, $3, $4) RETURNING id", task.Title, task.Description, task.Date, task.Completed).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *Repository) GetTask(id int) (*entity.Task, error) {
	var task entity.Task
	err := r.QueryRow("SELECT id, title, description, date, completed FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Date, &task.Completed)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *Repository) UpdateTask(id int, task *entity.Task) error {
	err := r.QueryRow("UPDATE tasks SET title=$1, description=$2, date=$3, completed=$4 WHERE id = $5", task.Title, task.Description, task.Date, task.Completed, id).Scan()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTask(id int) error {
	err := r.QueryRow("DELETE FROM tasks WHERE id = $1", id).Scan()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetTaskList(offset int, completed string, pagesize int, date string) ([]*entity.Task, error) {
	query := `SELECT id, title, description, date, completed FROM tasks`
	var args []interface{}
	whereClause := ""

	var count int
	if completed != "" {
		count++
		whereClause = fmt.Sprintf(" WHERE completed = $%d", count)
		args = append(args, completed)
	} else {
		whereClause = " WHERE (completed = 'true' OR completed = 'false')"
	}

	if date != "" {
		if whereClause != "" {
			whereClause += " AND "
		}
		whereClause += fmt.Sprintf("date = $%d", count+1)
		count += 1
		args = append(args, date)
	}

	query += whereClause + fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", count+1, count+2)

	fmt.Println(args)
	fmt.Println(query)

	rows, err := r.Query(query, append(args, pagesize, offset)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Date, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}
