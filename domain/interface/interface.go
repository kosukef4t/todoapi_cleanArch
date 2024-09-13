// domain/repository/todo_repository.go
package repository

import "myproject/domain/entity"

type IsToDoRepository interface {
	Get(title string, body string, startDate string, endDate string) ([]*entity.ToDo, error)
	GetByID(id string) (*entity.ToDo, error)
	Save(todo *entity.ToDo) error
	Update(id string, title string, body string, duedate string, completedAt string) (*entity.ToDo, error)
	IsCompleted(id string) (*entity.ToDo, error)
	Duplicate(id string) (*entity.ToDo, error)
	Delete(id string) error
}

type IsStaffRepository interface {
	Get(name, role string) ([]*entity.Staff, error)
	GetByStaff_ID(id string) (*entity.Staff, error)
	Save(staff *entity.Staff) error
	Update(id, name, role string) (*entity.Staff, error)
	Delete(id string) error
}

type IsToDoStaffRepository interface {
	Assign(todo_staff *entity.ToDo_Staff) error
}
