package entity

import "time"

type ToDo_Staff struct {
	id        string
	todo_id   string
	staff_id  string
	createdAt time.Time
	updatedAt time.Time
}

//インスタンスを作成する、これによりデータ構造を操作しなくて済む
func NewToDo_Staff(id, todo_id, staff_id string, createdAt, updatedAt time.Time) ToDo_Staff {
	return ToDo_Staff{
		id:        id,
		todo_id:   todo_id,
		staff_id:  staff_id,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (t *ToDo_Staff) ID() string {
	return t.id
}

func (t *ToDo_Staff) ToDo_ID() string {
	return t.todo_id
}

func (t *ToDo_Staff) Staff_ID() string {
	return t.staff_id
}

func (t *ToDo_Staff) CreatedAt() time.Time {
	return t.createdAt
}

func (t *ToDo_Staff) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *ToDo_Staff) UpdateUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}
