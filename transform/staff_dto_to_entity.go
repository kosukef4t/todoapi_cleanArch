package transform

import (
	"myproject/domain/entity"
	"myproject/dto"
)

func DtoToEntity_Staff(d *dto.Staff) entity.Staff {
	return entity.NewStaff(
		d.ID,
		d.Name,
		d.Role,
		d.CreatedAt,
		d.UpdatedAt,
	)
}

func EntityToDTO_Staff(entity *entity.Staff, fields ...string) dto.Staff {
	staff := dto.Staff{}

	// 全てのフィールドを処理
	staff.ID = entity.ID()
	staff.Name = entity.Name()
	staff.Role = entity.Role()
	staff.CreatedAt = entity.CreatedAt()
	staff.UpdatedAt = entity.UpdatedAt()

	return staff
}

func EntityToDTO_Staffs(entities []*entity.Staff, fields ...string) []dto.Staff {
	staffDTOs := make([]dto.Staff, len(entities))
	for i, e := range entities {
		staff := dto.Staff{}

		// 全てのフィールドを処理
		staff.ID = e.ID()
		staff.Name = e.Name()
		staff.Role = e.Role()
		staff.CreatedAt = e.CreatedAt()
		staff.UpdatedAt = e.UpdatedAt()

		staffDTOs[i] = staff
	}
	return staffDTOs
}
