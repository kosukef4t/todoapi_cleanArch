package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"myproject/domain/entity"
	repository "myproject/domain/interface"
	"myproject/infrastructure/database/models"
	"myproject/transform"
)

type ToDoStaffRepository struct {
	DB *sql.DB
}

func NewToDoStaffRepository(db *sql.DB) repository.IsToDoStaffRepository {
	return &ToDoStaffRepository{DB: db}
}

func (r *ToDoStaffRepository) GetByID(id string) (*entity.ToDo_Staff, error) {
	sqlQuery := "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	todo_staff := new(models.ToDo_Staff)
	if err := row.Scan(&todo_staff.ID, &todo_staff.ToDo_ID, &todo_staff.Staff_ID, &todo_staff.CreatedAt, &todo_staff.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			// IDに該当するデータが見つからない場合のエラーハンドリング
			return nil, fmt.Errorf("todo with ID %s not found", id)
		}
		return nil, err
	}

	entity := transform.ModelToEntity_ToDoStaff(todo_staff)
	return &entity, nil
}

func (r *ToDoStaffRepository) Get(todo_id, staff_id string) ([]*entity.ToDo_Staff, error) {
	args := []interface{}{}
	sqlQuery := "SELECT id, todo_id, staff_id, createdAt, updatedAt FROM todo_staff"

	if todo_id != "" {
		sqlQuery = "SELECT id, todo_id, staff_id, createdAt, updatedAt FROM todo_staff WHERE todo_id LIKE ? ORDER BY createdAt DESC"
		args = append(args, "%"+todo_id+"%")
	}

	if staff_id != "" {
		sqlQuery = "SELECT id, todo_id, staff_id, createdAt, updatedAt FROM todo_staff WHERE staff_id LIKE ? ORDER BY createdAt DESC"
		args = append(args, "%"+staff_id+"%")
	}

	rows, err := r.DB.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todo_staffs []*models.ToDo_Staff
	for rows.Next() {
		todo_staff := new(models.ToDo_Staff)
		if err := rows.Scan(&todo_staff.ID, &todo_staff.ToDo_ID, &todo_staff.Staff_ID, &todo_staff.CreatedAt, &todo_staff.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("todo_staff not found")
			}
			return nil, err
		}
		todo_staffs = append(todo_staffs, todo_staff)
	}

	entities := transform.ModelToEntity_ToDoStaffs(todo_staffs)
	entityPtrs := make([]*entity.ToDo_Staff, len(entities))
	for i := range entities {
		entityPtrs[i] = &entities[i] // ポインタをスライスに追加
	}

	return entityPtrs, nil

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

func (r *ToDoStaffRepository) Update(id, todo_id, staff_id string) (*entity.ToDo_Staff, error) {
	updateFields := []string{}
	args := []interface{}{}

	if todo_id != "" {
		updateFields = append(updateFields, "todo_id=?")
		args = append(args, todo_id)
	}

	if staff_id != "" {
		updateFields = append(updateFields, "staff_id=?")
		args = append(args, staff_id)
	}

	query := "UPDATE todo_staff SET " + strings.Join(updateFields, ",") + " WHERE id = ?"
	args = append(args, id)
	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update todo with ID %s: %w", id, err)
	}

	sqlQuery := "SELECT id, todo_id, staff_id, createdAt, updatedAt FROM todo_staff WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	todo_staff := new(models.ToDo_Staff)
	if err := row.Scan(&todo_staff.ID, &todo_staff.ToDo_ID, &todo_staff.Staff_ID, &todo_staff.CreatedAt, &todo_staff.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("staff with ID %s not found", id)
		}
		return nil, err
	}

	entity := transform.ModelToEntity_ToDoStaff(todo_staff)
	return &entity, nil
}

func (r *ToDoStaffRepository) Delete(id string) error {
	query := "DELETE FROM staffs WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
