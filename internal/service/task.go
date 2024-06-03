package service

import (
	"errors"
	"todo-list/internal/entity"
)

var ErrInvalidData = errors.New("invalid data")

func (s *Service) CreateTask(task *entity.Task) (int64, error) {
	if task.Title == "" || task.Date.IsZero() {
		return -1, ErrInvalidData
	}

	return s.TaskRepository.InsertTask(task)
}

func (s *Service) GetTask(id int) (*entity.Task, error) {
	if id <= 0 {
		return nil, ErrInvalidData
	}

	return s.TaskRepository.GetTask(id)
}

func (s *Service) UpdateTask(id int, task *entity.Task) error {
	if task.Title == "" || task.Date.IsZero() || id <= 0 {
		return ErrInvalidData
	}

	return s.TaskRepository.UpdateTask(id, task)
}

func (s *Service) DeleteTask(id int) error {
	if id <= 0 {
		return ErrInvalidData
	}

	return s.TaskRepository.DeleteTask(id)
}

func (s *Service) GetTaskList(offset int, completed string, pagesize int, date string) ([]*entity.Task, error) {
	if offset < 0 || pagesize <= 0 {
		return nil, ErrInvalidData
	}

	return s.TaskRepository.GetTaskList(offset, completed, pagesize, date)
}
