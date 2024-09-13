// infrastructure/repository/todo_repository.go
package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"myproject/domain/entity"
	repository "myproject/domain/interface"
	"myproject/infrastructure/database/models"
	"myproject/transform"

	"github.com/lucsky/cuid"
)

type ToDoRepository struct {
	DB *sql.DB
}

func NewToDoRepository(db *sql.DB) repository.IsToDoRepository {
	return &ToDoRepository{DB: db}
}

func (r *ToDoRepository) GetByID(id string) (*entity.ToDo, error) {
	sqlQuery := "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	todo := new(models.ToDo)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompletedAt, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			// IDに該当するデータが見つからない場合のエラーハンドリング
			return nil, fmt.Errorf("todo with ID %s not found", id)
		}
		return nil, err
	}

	entity := transform.Model_To_Entity(todo)
	return &entity, nil
}

func (r *ToDoRepository) Get(title string, body string, startDate string, endDate string) ([]*entity.ToDo, error) {
	args := []interface{}{}
	sqlQuery := "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos"

	if title != "" {
		sqlQuery = "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos WHERE title LIKE ? ORDER BY createdAt DESC"
		args = append(args, "%"+title+"%")
	}

	if body != "" {
		sqlQuery = "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos WHERE body LIKE ? ORDER BY createdAt DESC"
		args = append(args, "%"+body+"%")
	}

	if startDate != "" && endDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, fmt.Errorf("無効な開始日フォーマット: %v", err)
		}

		parsedEndDate, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			return nil, fmt.Errorf("無効な終了日フォーマット: %v", err)
		}
		sqlQuery = "SELECT id, title, body, duedate FROM todos WHERE duedate BETWEEN ? AND ?"
		args = append(args, parsedStartDate, parsedEndDate)
	}

	rows, err := r.DB.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*models.ToDo
	for rows.Next() {
		todo := new(models.ToDo)
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompletedAt, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("todo not found")
			}
			return nil, err
		}
		todos = append(todos, todo)
	}

	entities := transform.ModelToEntity(todos)
	entityPtrs := make([]*entity.ToDo, len(entities))
	for i := range entities {
		entityPtrs[i] = &entities[i] // ポインタをスライスに追加
	}

	return entityPtrs, nil

}

func (r *ToDoRepository) Save(todo *entity.ToDo) error {
	todoModel := transform.EntityToModel(todo)
	query := "INSERT INTO todos (id, title, body, duedate) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, todoModel.ID, todoModel.Title, todoModel.Body, todoModel.DueDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *ToDoRepository) Update(id, title, body, duedate, completedAt string) (*entity.ToDo, error) {
	updateFields := []string{}
	args := []interface{}{}

	if title != "" {
		updateFields = append(updateFields, "title=?")
		args = append(args, title)
	}

	if body != "" {
		updateFields = append(updateFields, "body=?")
		args = append(args, body)
	}

	if duedate != "" {
		updateFields = append(updateFields, "duedate = ?")
		args = append(args, duedate)
	}

	if completedAt != "" {
		updateFields = append(updateFields, "completedAt = ?")
		args = append(args, completedAt)
	}

	query := "UPDATE todos SET " + strings.Join(updateFields, ",") + " WHERE id = ?"
	args = append(args, id)
	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update todo with ID %s: %w", id, err)
	}

	sqlQuery := "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	todo := new(models.ToDo)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompletedAt, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			// IDに該当するデータが見つからない場合のエラーハンドリング
			return nil, fmt.Errorf("todo with ID %s not found", id)
		}
		return nil, err
	}

	entity := transform.Model_To_Entity(todo)
	return &entity, nil
}

func (r *ToDoRepository) IsCompleted(id string) (*entity.ToDo, error) {
	sqlQuery := "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	todo := new(models.ToDo)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.DueDate, &todo.CompletedAt, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			// IDに該当するデータが見つからない場合のエラーハンドリング
			return nil, fmt.Errorf("todo with ID %s not found", id)
		}
		return nil, err
	}

	entity := transform.Model_To_Entity(todo)
	return &entity, nil
}

func (r *ToDoRepository) Duplicate(id string) (*entity.ToDo, error) {
	sqlQuery := "SELECT id, title, body, duedate, completedAt, createdAt, updatedAt FROM todos WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	todoModel := new(models.ToDo)
	if err := row.Scan(&todoModel.ID, &todoModel.Title, &todoModel.Body, &todoModel.DueDate, &todoModel.CompletedAt, &todoModel.CreatedAt, &todoModel.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			// IDに該当するデータが見つからない場合のエラーハンドリング
			return nil, fmt.Errorf("todo with ID %s not found", id)
		}
		return nil, err
	}

	todoModel.ID = cuid.New()
	todoModel.Title = "{" + todoModel.Title + "}コピー"
	todoModel.Body = ""
	todoModel.DueDate = sql.NullTime{}
	todoModel.CompletedAt = sql.NullTime{}
	todoModel.CreatedAt = time.Now()
	todoModel.UpdatedAt = time.Now()

	query := "INSERT INTO todos (id, title, body, duedate, completedAt, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, todoModel.ID, todoModel.Title, todoModel.Body, todoModel.DueDate, todoModel.CompletedAt, todoModel.CreatedAt, todoModel.UpdatedAt)
	if err != nil {
		return nil, err
	}

	entity := transform.Model_To_Entity(todoModel)
	return &entity, nil
}

func (r *ToDoRepository) Delete(id string) error {
	query := "DELETE FROM todos WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
