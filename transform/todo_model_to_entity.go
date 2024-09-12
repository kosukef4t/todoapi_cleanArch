package transform

import (
	"database/sql"
	"myproject/domain/entity"
	"myproject/infrastructure/database/models"
	"time"
)

func ModelToEntity(models []*models.ToDo) []entity.ToDo {
	entities := make([]entity.ToDo, len(models))
	for i, m := range models {
		entities[i] = entity.NewToDo(
			m.ID,
			m.Title,
			m.Body,
			func() time.Time {
				if m.DueDate.Valid {
					return m.DueDate.Time
				}
				return time.Time{}
			}(),
			func() time.Time {
				if m.CompletedAt.Valid {
					return m.CompletedAt.Time
				}
				return time.Time{}
			}(),
			m.CreatedAt,
			m.UpdatedAt,
		)
	}
	return entities
}

func Model_To_Entity(model *models.ToDo) entity.ToDo {
	return entity.NewToDo(
		model.ID,
		model.Title,
		model.Body,
		func() time.Time {
			if model.DueDate.Valid {
				return model.DueDate.Time
			}
			return time.Time{}
		}(),
		func() time.Time {
			if model.CompletedAt.Valid {
				return model.CompletedAt.Time
			}
			return time.Time{}
		}(),
		model.CreatedAt,
		model.UpdatedAt,
	)
}

func EntityToModel(e *entity.ToDo, fields ...string) models.ToDo {
	todo := models.ToDo{}

	todo.ID = e.ID()
	todo.Title = e.Title()
	todo.Body = e.Body()

	if e.DueDate() != (time.Time{}) {
		todo.DueDate = sql.NullTime{
			Time:  e.DueDate(),
			Valid: true,
		}
	} else {
		todo.DueDate = sql.NullTime{Valid: false}
	}

	if e.CompletedAt() != (time.Time{}) {
		todo.CompletedAt = sql.NullTime{
			Time:  e.CompletedAt(),
			Valid: true,
		}
	} else {
		todo.CompletedAt = sql.NullTime{Valid: false}
	}

	todo.CreatedAt = e.CreatedAt()
	todo.UpdatedAt = e.UpdatedAt()

	return todo
}
