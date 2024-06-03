package repository

import (
	"testing"
	"time"
	"todo-list/internal/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := &Repository{db}

	task := &entity.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Date:        time.Now(),
		Completed:   false,
	}

	mock.ExpectQuery("INSERT INTO tasks").
		WithArgs(task.Title, task.Description, task.Date, task.Completed).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := repo.InsertTask(task)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := &Repository{db}

	task := &entity.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Date:        time.Now(),
		Completed:   false,
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "date", "completed"}).
		AddRow(task.ID, task.Title, task.Description, task.Date, task.Completed)

	mock.ExpectQuery("SELECT id, title, description, date, completed FROM tasks WHERE id = \\$1").
		WithArgs(task.ID).
		WillReturnRows(rows)

	result, err := repo.GetTask(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := &Repository{db}

	task := &entity.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Date:        time.Now(),
		Completed:   true,
	}

	mock.ExpectQuery("UPDATE tasks SET title=\\$1, description=\\$2, date=\\$3, completed=\\$4 WHERE id = \\$5").
		WithArgs(task.Title, task.Description, task.Date, task.Completed, 1).
		WillReturnRows(sqlmock.NewRows([]string{}))

	err = repo.UpdateTask(1, task)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := &Repository{db}

	mock.ExpectQuery("DELETE FROM tasks WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{}))

	err = repo.DeleteTask(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTaskList(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := &Repository{db}

	tasks := []*entity.Task{
		{
			ID:          1,
			Title:       "Test Task 1",
			Description: "Description 1",
			Date:        time.Now(),
			Completed:   false,
		},
		{
			ID:          2,
			Title:       "Test Task 2",
			Description: "Description 2",
			Date:        time.Now(),
			Completed:   true,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "date", "completed"}).
		AddRow(tasks[0].ID, tasks[0].Title, tasks[0].Description, tasks[0].Date, tasks[0].Completed).
		AddRow(tasks[1].ID, tasks[1].Title, tasks[1].Description, tasks[1].Date, tasks[1].Completed)

	mock.ExpectQuery("SELECT id, title, description, date, completed FROM tasks WHERE \\(completed = 'true' OR completed = 'false'\\) ORDER BY id LIMIT \\$1 OFFSET \\$2").
		WithArgs(10, 0).
		WillReturnRows(rows)

	result, err := repo.GetTaskList(0, "", 10, "")
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}
