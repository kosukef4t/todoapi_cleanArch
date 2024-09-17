package transform

import (
	"myproject/domain/entity"
	"myproject/infrastructure/database/models"
)

func EntityToModel_ToDoStaff(e *entity.ToDo_Staff, fields ...string) models.ToDo_Staff {
	todo_staff := models.ToDo_Staff{}

	todo_staff.ID = e.ID()
	todo_staff.ToDo_ID = e.ToDo_ID()
	todo_staff.Staff_ID = e.Staff_ID()
	todo_staff.CreatedAt = e.CreatedAt()
	todo_staff.UpdatedAt = e.UpdatedAt()

	return todo_staff
}

func ModelToEntity_ToDoStaff(model *models.ToDo_Staff) entity.ToDo_Staff {
	return entity.NewToDo_Staff(
		model.ID,
		model.ToDo_ID,
		model.Staff_ID,
		model.CreatedAt,
		model.UpdatedAt,
	)
}

func ModelToEntity_ToDoStaffs(models []*models.ToDo_Staff) []entity.ToDo_Staff {
	entities := make([]entity.ToDo_Staff, len(models))
	for i, m := range models {
		entities[i] = entity.NewToDo_Staff(
			m.ID,
			m.ToDo_ID,
			m.Staff_ID,
			m.CreatedAt,
			m.UpdatedAt,
		)
	}
	return entities
}
