// application/services/todo_service.go
package services

import (
	"myproject/domain/entity"
	"myproject/domain/repository"
)

type Service struct {
	RepoToDo       repository.IsToDoRepository
	RepoStaff      repository.IsStaffRepository
	RepoToDo_Staff repository.IsToDoStaffRepository
}

func NewService(
	repoToDo repository.IsToDoRepository,
	repoStaff repository.IsStaffRepository,
	repoToDoStaff repository.IsToDoStaffRepository,
) *Service {
	return &Service{
		RepoToDo:       repoToDo,
		RepoStaff:      repoStaff,
		RepoToDo_Staff: repoToDoStaff,
	}
}

// 以下からtodoに関するビジネスロジック
func (s *Service) GetByID(id string) (*entity.ToDo, error) {
	return s.RepoToDo.GetByID(id)
}

func (s *Service) GetToDo(title, body, startDate, endDate string) ([]*entity.ToDo, error) {
	return s.RepoToDo.Get(title, body, startDate, endDate)
}

func (s *Service) CreateToDo(todo *entity.ToDo) error {
	return s.RepoToDo.Save(todo)
}

func (s *Service) UpdateTodo(id, title, body, duedate, completedAt string) (*entity.ToDo, error) {
	return s.RepoToDo.Update(id, title, body, duedate, completedAt)
}

func (s *Service) CompletedToDo(id string) (*entity.ToDo, error) {
	return s.RepoToDo.IsCompleted(id)
}

func (s *Service) DuplicateToDo(id string) (*entity.ToDo, error) {
	return s.RepoToDo.Duplicate(id)
}

func (s *Service) DeleteToDoByID(id string) error {
	return s.RepoToDo.Delete(id)
}

// 以下よりStuffに関するビジネスロジック
func (s *Service) GetStaff(name, role string) ([]*entity.Staff, error) {
	return s.RepoStaff.Get(name, role)
}

func (s *Service) GetByStaff_ID(id string) (*entity.Staff, error) {
	return s.RepoStaff.GetByStaff_ID(id)
}

func (s *Service) CreateStaff(staff *entity.Staff) error {
	return s.RepoStaff.Save(staff)
}

func (s *Service) UpdateStaff(id, name, role string) (*entity.Staff, error) {
	return s.RepoStaff.Update(id, name, role)
}

func (s *Service) DeleteStaff(id string) error {
	return s.RepoStaff.Delete(id)
}

// 以下よりToDo_Stuffに関するビジネスロジック
func (s *Service) Assigntodo(todo_staff *entity.ToDo_Staff) error {
	return s.RepoToDo_Staff.Assign(todo_staff)
}
