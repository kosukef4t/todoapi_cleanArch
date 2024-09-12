// domain/repository/todo_repository.go
package repository

import entity "myproject/domain/entity/auth"

type IsUserRepository interface {
	Create(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
}
