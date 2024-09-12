package auth

import (
	"database/sql"
	entity "myproject/domain/entity/auth"
	repository "myproject/domain/repository/auth"
	auth_models "myproject/infrastructure/auth/models"
	auth_transform "myproject/transform/auth"
)

// UserRepositoryImplは実際のデータベース操作を行う実装
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IsUserRepository {
	return &UserRepository{db: db}
}

// Create関数：新しいユーザーをデータベースに追加
func (r *UserRepository) Create(user *entity.User) error {
	user_model := auth_transform.EntityToModel_Auth(user)
	_, err := r.db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", user_model.ID, user_model.Username, user_model.Password)
	return err
}

// FindByUsername関数：ユーザー名でユーザーを検索
func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
	user := new(auth_models.User)
	err := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	entity_user := auth_transform.ModelToEntity_Auth(user)
	return &entity_user, nil
}
