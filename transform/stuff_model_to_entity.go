package transform

import (
	"myproject/domain/entity"
	"myproject/infrastructure/database/models"
)

func EntityToModel_Staff(e *entity.Staff, fields ...string) models.Staff {
	staff := models.Staff{}

	staff.ID = e.ID()
	staff.Name = e.Name()
	staff.Role = e.Role()
	staff.CreatedAt = e.CreatedAt()
	staff.UpdatedAt = e.UpdatedAt()

	return staff
}

func ModelToEntity_Staff(model *models.Staff) entity.Staff {
	return entity.NewStaff(
		model.ID,
		model.Name,
		model.Role,
		model.CreatedAt,
		model.UpdatedAt,
	)
}

func ModelToEntity_Staffs(models []*models.Staff) []entity.Staff {
	entities := make([]entity.Staff, len(models))
	for i, m := range models {
		entities[i] = entity.NewStaff(
			m.ID,
			m.Name,
			m.Role,
			m.CreatedAt,
			m.UpdatedAt,
		)
	}
	return entities
}
