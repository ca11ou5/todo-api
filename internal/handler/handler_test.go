package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(task *entity.Task) (int64, error) {
	args := m.Called(task)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTaskService) GetTask(id int) (*entity.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(id int, task *entity.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskService) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskService) GetTaskList(offset int, completed string, pagesize int, date string) ([]*entity.Task, error) {
	args := m.Called(offset, completed, pagesize, date)
	return args.Get(0).([]*entity.Task), args.Error(1)
}

func setupRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	r.POST("task", h.CreateTask)
	r.GET("task/:id", h.GetTask)
	r.PUT("task/:id", h.UpdateTask)
	r.DELETE("task/:id", h.DeleteTask)
	r.GET("task", h.GetTaskList)

	return r
}

func TestCreateTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewHandler(mockService)
	router := setupRouter(handler)

	task := &entity.Task{
		Title:       "Test Task",
		Description: "Test Description",
	}
	mockService.On("CreateTask", task).Return(int64(1), nil)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/task", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewHandler(mockService)
	router := setupRouter(handler)

	task := &entity.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
	}
	mockService.On("GetTask", 1).Return(task, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/task/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseTask entity.Task
	_ = json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.Equal(t, task, &responseTask)
	mockService.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewHandler(mockService)
	router := setupRouter(handler)

	task := &entity.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
	}
	mockService.On("UpdateTask", 1, task).Return(nil)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/task/1", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewHandler(mockService)
	router := setupRouter(handler)

	mockService.On("DeleteTask", 1).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/task/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetTaskList(t *testing.T) {
	mockService := new(MockTaskService)
	handler := NewHandler(mockService)
	router := setupRouter(handler)

	tasks := []*entity.Task{
		{
			ID:          1,
			Title:       "Test Task 1",
			Description: "Test Description 1",
		},
		{
			ID:          2,
			Title:       "Test Task 2",
			Description: "Test Description 2",
		},
	}
	mockService.On("GetTaskList", 0, "", 10, "").Return(tasks, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/task?page=1&pageSize=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseTasks []*entity.Task
	_ = json.Unmarshal(w.Body.Bytes(), &responseTasks)
	assert.Equal(t, tasks, responseTasks)
	mockService.AssertExpectations(t)
}
