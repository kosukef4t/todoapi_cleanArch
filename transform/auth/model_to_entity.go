package auth_transform

import (
	entity "myproject/domain/entity/auth"
	auth_models "myproject/infrastructure/auth/models"
)

func EntityToModel_Auth(e *entity.User, fields ...string) auth_models.User {
	auth := auth_models.User{}

	auth.ID = e.ID()
	auth.Username = e.Username()
	auth.Password = e.Password()
	auth.CreatedAt = e.CreatedAt()
	auth.UpdatedAt = e.UpdatedAt()

	return auth
}

func ModelToEntity_Auth(model *auth_models.User) entity.User {
	return entity.NewUser(
		model.ID,
		model.Username,
		model.Password,
		model.CreatedAt,
		model.UpdatedAt,
	)
}
