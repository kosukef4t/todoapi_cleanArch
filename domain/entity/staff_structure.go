package entity

import "time"

type Staff struct {
	id        string
	name      string
	role      string
	createdAt time.Time
	updatedAt time.Time
}

//インスタンスを作成する、これによりデータ構造を操作しなくて済む
func NewStaff(id, name, role string, createdAt, updatedAt time.Time) Staff {
	return Staff{
		id:        id,
		name:      name,
		role:      role,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (t *Staff) ID() string {
	return t.id
}

func (t *Staff) Name() string {
	return t.name
}

func (t *Staff) Role() string {
	return t.role
}

func (t *Staff) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Staff) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Staff) UpdateName(name string) {
	t.name = name
}

func (t *Staff) UpdateRole(role string) {
	t.role = role
}

func (t *Staff) UpdateUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}
