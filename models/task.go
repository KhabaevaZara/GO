package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"oneof=todo in_progress done"`
	CreatedAt   time.Time `json:"created_at"`
}
