package transform

import (
	"database/sql"
	"myproject/domain/entity"
	"myproject/dto"
	"time"
)

func DtoToEntity(d *dto.ToDo) entity.ToDo {
	return entity.NewToDo(
		d.ID,
		d.Title,
		d.Body,
		func() time.Time {
			if d.DueDate.Valid {
				return d.DueDate.Time
			}
			return time.Time{}
		}(),
		func() time.Time {
			if d.CompletedAt.Valid {
				return d.CompletedAt.Time
			}
			return time.Time{}
		}(),
		d.CreatedAt,
		d.UpdatedAt,
	)
}

func EntityToDTO(entities []*entity.ToDo, fields ...string) []dto.ToDo {
	todoDTOs := make([]dto.ToDo, len(entities))
	for i, e := range entities {
		todo := dto.ToDo{}

		// 全てのフィールドを処理
		todo.ID = e.ID()
		todo.Title = e.Title()
		todo.Body = e.Body()

		// DueDateの処理
		if e.DueDate() != (time.Time{}) {
			todo.DueDate = sql.NullTime{
				Time:  e.DueDate(),
				Valid: true,
			}
		} else {
			todo.DueDate = sql.NullTime{Valid: false}
		}

		// CompletedAtの処理
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

		todoDTOs[i] = todo
	}
	return todoDTOs
}

func Entity_To_DTO(entity *entity.ToDo, fields ...string) dto.ToDo {
	todo := dto.ToDo{}

	// 全てのフィールドを処理
	todo.ID = entity.ID()
	todo.Title = entity.Title()
	todo.Body = entity.Body()

	// DueDateの処理
	if entity.DueDate() != (time.Time{}) {
		todo.DueDate = sql.NullTime{
			Time:  entity.DueDate(),
			Valid: true,
		}
	} else {
		todo.DueDate = sql.NullTime{Valid: false}
	}

	// CompletedAtの処理
	if entity.CompletedAt() != (time.Time{}) {
		todo.CompletedAt = sql.NullTime{
			Time:  entity.CompletedAt(),
			Valid: true,
		}
	} else {
		todo.CompletedAt = sql.NullTime{Valid: false}
	}

	todo.CreatedAt = entity.CreatedAt()
	todo.UpdatedAt = entity.UpdatedAt()

	return todo
}
