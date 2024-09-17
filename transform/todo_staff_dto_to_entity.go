package transform

import (
	"myproject/domain/entity"
	"myproject/dto"
)

func DtoToEntity_ToDoStaff(d *dto.ToDo_Staff) entity.ToDo_Staff {
	return entity.NewToDo_Staff(
		d.ID,
		d.ToDo_ID,
		d.Staff_ID,
		d.CreatedAt,
		d.UpdatedAt,
	)
}

func EntityToDto_ToDoStaff(entity *entity.ToDo_Staff, fields ...string) dto.ToDo_Staff {
	ToDo_Staff := dto.ToDo_Staff{}

	// 全てのフィールドを処理
	ToDo_Staff.ID = entity.ID()
	ToDo_Staff.ToDo_ID = entity.ToDo_ID()
	ToDo_Staff.Staff_ID = entity.Staff_ID()
	ToDo_Staff.CreatedAt = entity.CreatedAt()
	ToDo_Staff.UpdatedAt = entity.UpdatedAt()

	return ToDo_Staff
}

func EntityToDto_ToDoStaffs(entities []*entity.ToDo_Staff, fields ...string) []dto.ToDo_Staff {
	todo_staffDTOs := make([]dto.ToDo_Staff, len(entities))
	for i, e := range entities {
		todo_staff := dto.ToDo_Staff{}
		// 全てのフィールドを処理
		todo_staff.ID = e.ID()
		todo_staff.ToDo_ID = e.ToDo_ID()
		todo_staff.Staff_ID = e.Staff_ID()
		todo_staff.CreatedAt = e.CreatedAt()
		todo_staff.UpdatedAt = e.UpdatedAt()

		todo_staffDTOs[i] = todo_staff
	}
	return todo_staffDTOs
}
