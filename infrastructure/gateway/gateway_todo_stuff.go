package repository

import (
	"database/sql"

	"myproject/domain/entity"
	"myproject/domain/repository"
	"myproject/transform"
)

type ToDoStaffRepository struct {
	DB *sql.DB
}

func NewToDoStaffRepository(db *sql.DB) repository.IsToDoStaffRepository {
	return &ToDoStaffRepository{DB: db}
}

func (r *ToDoStaffRepository) Assign(todo_staff *entity.ToDo_Staff) error {
	todo_staffModel := transform.EntityToModel_ToDoStaff(todo_staff)
	query := "INSERT INTO todo_staff (id, todo_id, staff_id) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, todo_staffModel.ID, todo_staffModel.ToDo_ID, todo_staffModel.Staff_ID)
	if err != nil {
		return err
	}

	return nil
}
