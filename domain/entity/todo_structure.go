package entity

import "time"

type ToDo struct {
	id          string
	title       string
	body        string
	dueDate     time.Time
	completedAt time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

//インスタンスを作成する、これによりデータ構造を操作しなくて済む
func NewToDo(id, title, body string, dueDate, completedAt, createdAt, updatedAt time.Time) ToDo {
	return ToDo{
		id:          id,
		title:       title,
		body:        body,
		dueDate:     dueDate,
		completedAt: completedAt,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (t *ToDo) ID() string {
	return t.id
}

func (t *ToDo) Title() string {
	return t.title
}

func (t *ToDo) Body() string {
	return t.body
}

func (t *ToDo) DueDate() time.Time {
	return t.dueDate
}

func (t *ToDo) CompletedAt() time.Time {
	return t.completedAt
}

func (t *ToDo) CreatedAt() time.Time {
	return t.createdAt
}

func (t *ToDo) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *ToDo) UpdateTitle(title string) {
	t.title = title
}

func (t *ToDo) UpdateBody(body string) {
	t.body = body
}

func (t *ToDo) UpdateDueDate(dueDate time.Time) {
	t.dueDate = dueDate
}

func (t *ToDo) UpdateCompletedAt(completedAt time.Time) {
	t.completedAt = completedAt
}

func (t *ToDo) UpdateUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}
