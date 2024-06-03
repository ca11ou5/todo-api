package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/internal/entity"
	"todo-list/internal/service"
)

func NewHandler(service TaskService) *Handler {
	return &Handler{service}
}

type Handler struct {
	TaskService
}

type TaskService interface {
	CreateTask(task *entity.Task) (int64, error)
	GetTask(id int) (*entity.Task, error)
	UpdateTask(id int, task *entity.Task) error
	DeleteTask(id int) error
	GetTaskList(offset int, completed string, pagesize int, date string) ([]*entity.Task, error)
}

// CreateTask godoc
//
//	@Summary		Create a task
//	@Description	Create a new task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		entity.Task	true	"Task"
//	@Success		201		{object}	map[string]int64
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/task [post]
func (h *Handler) CreateTask(ctx *gin.Context) {
	var task entity.Task

	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.TaskService.CreateTask(&task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetTask godoc
//
//	@Summary		Get a task
//	@Description	Get a task by ID
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	entity.Task
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/task/{id} [get]
func (h *Handler) GetTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.TaskService.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// UpdateTask godoc
//
//	@Summary		Update a task
//	@Description	Update a task by ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Task ID"
//	@Param			task	body		entity.Task	true	"Task"
//	@Success		200		{object}	map[string]int
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/task/{id} [put]
func (h *Handler) UpdateTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task entity.Task
	err = ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.TaskService.UpdateTask(id, &task)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// DeleteTask godoc
//
//	@Summary		Delete a task
//	@Description	Delete a task by ID
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	map[string]int
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/task/{id} [delete]
func (h *Handler) DeleteTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.TaskService.DeleteTask(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// GetTaskList godoc
//
//	@Summary		Get task list
//	@Description	Get a list of tasks with pagination
//	@Tags			tasks
//	@Produce		json
//	@Param			page		query		int		false	"Page number"				default(1)
//	@Param			pageSize	query		int		false	"Number of tasks per page"	default(10)
//	@Param			completed	query		string	false	"Filter by completion status"
//	@Param			date		query		string	false	"Filter by date"
//	@Success		200			{array}		entity.Task
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/task [get]
func (h *Handler) GetTaskList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	completed := ctx.DefaultQuery("completed", "")
	date := ctx.DefaultQuery("date", "")

	offset := (page - 1) * pageSize

	tasks, err := h.TaskService.GetTaskList(offset, completed, pageSize, date)
	if err != nil && !errors.Is(err, service.ErrInvalidData) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if errors.Is(err, service.ErrInvalidData) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
