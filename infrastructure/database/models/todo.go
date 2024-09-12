package models

import (
	"database/sql"
	"time"
)

type ToDo struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Body        string       `json:"body"`
	DueDate     sql.NullTime `json:"duedate,omitempty"`
	CompletedAt sql.NullTime `json:"completeAt,omitempty"`
	CreatedAt   time.Time    `json:"createdAt,omitempty"`
	UpdatedAt   time.Time    `json:"updatedAt,omitempty"`
}
