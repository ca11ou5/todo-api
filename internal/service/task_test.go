package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"todo-list/internal/entity"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) InsertTask(task *entity.Task) (int64, error) {
	args := m.Called(task)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTaskRepository) GetTask(id int) (*entity.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(id int, task *entity.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTaskList(offset int, completed string, pagesize int, date string) ([]*entity.Task, error) {
	args := m.Called(offset, completed, pagesize, date)
	return args.Get(0).([]*entity.Task), args.Error(1)
}

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	task := &entity.Task{
		Title: "Test Task",
		Date:  time.Now(),
	}

	mockRepo.On("InsertTask", task).Return(int64(1), nil)

	id, err := service.CreateTask(task)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
	mockRepo.AssertExpectations(t)
}

func TestCreateTask_InvalidData(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	task := &entity.Task{
		Title: "",
		Date:  time.Time{},
	}

	id, err := service.CreateTask(task)
	assert.Error(t, err)
	assert.Equal(t, int64(-1), id)
	assert.Equal(t, ErrInvalidData, err)
}

func TestGetTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	task := &entity.Task{
		ID:    1,
		Title: "Test Task",
		Date:  time.Now(),
	}

	mockRepo.On("GetTask", 1).Return(task, nil)

	result, err := service.GetTask(1)
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTask_InvalidID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	result, err := service.GetTask(-1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrInvalidData, err)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	task := &entity.Task{
		Title: "Updated Task",
		Date:  time.Now(),
	}

	mockRepo.On("UpdateTask", 1, task).Return(nil)

	err := service.UpdateTask(1, task)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_InvalidData(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	task := &entity.Task{
		Title: "",
		Date:  time.Time{},
	}

	err := service.UpdateTask(-1, task)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidData, err)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	mockRepo.On("DeleteTask", 1).Return(nil)

	err := service.DeleteTask(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_InvalidID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	err := service.DeleteTask(-1)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidData, err)
}

func TestGetTaskList(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewService(mockRepo)

	tasks := []*entity.Task{
		{
			ID:    1,
			Title: "Test Task 1",
			Date:  time.Now(),
		},
		{
			ID:    2,
			Title: "Test Task 2",
			Date:  time.Now(),
		},
	}

	mockRepo.On("GetTaskList", 0, "", 10, "").Return(tasks, nil)

	result, err := service.GetTaskList(0, "", 10, "")
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}
