package dto

import (
	"time"
)

type ToDo_Staff struct {
	ID        string    `json:"id"`
	ToDo_ID   string    `json:"todo_id"`
	Staff_ID  string    `json:"staff_id"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
