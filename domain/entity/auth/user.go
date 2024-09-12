package auth_entity

import "time"

// Userエンティティの定義
type User struct {
	id        string
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

//インスタンスを作成する、これによりデータ構造を操作しなくて済む
func NewUser(id, username, password string, createdAt, updatedAt time.Time) User {
	return User{
		id:        id,
		username:  username,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (t *User) ID() string {
	return t.id
}

func (t *User) Username() string {
	return t.username
}

func (t *User) Password() string {
	return t.password
}

func (t *User) CreatedAt() time.Time {
	return t.createdAt
}

func (t *User) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *User) UpdateUsername(username string) {
	t.username = username
}

func (t *User) UpdateUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}
