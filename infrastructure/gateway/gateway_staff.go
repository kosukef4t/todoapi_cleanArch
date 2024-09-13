package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"myproject/domain/entity"
	repository "myproject/domain/interface"
	"myproject/infrastructure/database/models"
	"myproject/transform"
)

type StaffRepository struct {
	DB *sql.DB
}

func NewStaffRepository(db *sql.DB) repository.IsStaffRepository {
	return &StaffRepository{DB: db}
}

func (r *StaffRepository) Get(name, role string) ([]*entity.Staff, error) {
	args := []interface{}{}
	sqlQuery := "SELECT id, name, role, createdAt, updatedAt FROM staffs"

	if name != "" {
		sqlQuery = "SELECT id, name, role, createdAt, updatedAt FROM staffs WHERE name LIKE ? ORDER BY createdAt DESC"
		args = append(args, "%"+name+"%")
	}

	if role != "" {
		sqlQuery = "SELECT id, name, role, createdAt, updatedAt FROM staffs WHERE role LIKE ? ORDER BY createdAt DESC"
		args = append(args, "%"+role+"%")
	}

	rows, err := r.DB.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staffs []*models.Staff
	for rows.Next() {
		staff := new(models.Staff)
		if err := rows.Scan(&staff.ID, &staff.Name, &staff.Role, &staff.CreatedAt, &staff.UpdatedAt); err != nil {
			return nil, err
		}
		staffs = append(staffs, staff)
	}

	entities := transform.ModelToEntity_Staffs(staffs)
	entityPtrs := make([]*entity.Staff, len(entities))
	for i := range entities {
		entityPtrs[i] = &entities[i] // ポインタをスライスに追加
	}

	return entityPtrs, nil

}

func (r *StaffRepository) GetByStaff_ID(id string) (*entity.Staff, error) {
	sqlQuery := "SELECT id, name, role, createdAt, updatedAT FROM staffs WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	staff := new(models.Staff)
	if err := row.Scan(&staff.ID, &staff.Name, &staff.Role, &staff.CreatedAt, &staff.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			// IDが見つからなかった場合、nilを返す
			return nil, fmt.Errorf("todo with ID %s not found", id)
		}
		return nil, err
	}

	entity := transform.ModelToEntity_Staff(staff)
	return &entity, nil
}

func (r *StaffRepository) Save(staff *entity.Staff) error {
	staffModel := transform.EntityToModel_Staff(staff)
	query := "INSERT INTO staffs (id, name, role) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, staffModel.ID, staffModel.Name, staffModel.Role)
	if err != nil {
		return fmt.Errorf("failed to save staff: %w", err)
	}

	return nil
}

func (r *StaffRepository) Update(id, name, role string) (*entity.Staff, error) {
	updateFields := []string{}
	args := []interface{}{}

	if name != "" {
		updateFields = append(updateFields, "name=?")
		args = append(args, name)
	}

	if role != "" {
		updateFields = append(updateFields, "role=?")
		args = append(args, role)
	}

	query := "UPDATE staffs SET " + strings.Join(updateFields, ",") + " WHERE id = ?"
	args = append(args, id)
	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update todo with ID %s: %w", id, err)
	}

	sqlQuery := "SELECT id, name, role, createdAt, updatedAt FROM staffs WHERE id = ?"
	row := r.DB.QueryRow(sqlQuery, id)

	staff := new(models.Staff)
	if err := row.Scan(&staff.ID, &staff.Name, &staff.Role, &staff.CreatedAt, &staff.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("staff with ID %s not found", id)
		}
		return nil, err
	}

	entity := transform.ModelToEntity_Staff(staff)
	return &entity, nil
}

func (r *StaffRepository) Delete(id string) error {
	query := "DELETE FROM staffs WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
