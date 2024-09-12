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
