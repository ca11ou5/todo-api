package service

import "todo-list/internal/entity"

type Service struct {
	TaskRepository
}

func NewService(repo TaskRepository) *Service {
	return &Service{repo}
}

type TaskRepository interface {
	InsertTask(task *entity.Task) (int64, error)
	GetTask(id int) (*entity.Task, error)
	UpdateTask(id int, task *entity.Task) error
	DeleteTask(id int) error
	GetTaskList(offset int, completed string, pagesize int, date string) ([]*entity.Task, error)
}
