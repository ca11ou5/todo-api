package entity

import "time"

type Task struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Task title"`
	Description string    `json:"description" example:"Task description"`
	Date        time.Time `json:"date" example:"2020-01-01T00:00:00Z"`
	Completed   bool      `json:"completed" example:"true"`
}
