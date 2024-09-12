package auth_services

import (
	"errors"
	"fmt"
	auth_entity "myproject/domain/entity/auth"
	repository "myproject/domain/repository/auth"
	"time"

	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthService構造体：認証サービスを管理
type AuthService struct {
	UserRepo repository.IsUserRepository
}

// NewAuthService関数：ユーザーリポジトリを受け取って認証サービスを作成
func NewAuthService(userRepo repository.IsUserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) RegisterUser(username, password string) (string, error) {
	// ユーザー名の重複確認
	existingUser, _ := s.UserRepo.FindByUsername(username)
	if existingUser != nil {
		return "", fmt.Errorf("user already exists") // ユーザー名がすでに存在する場合のエラー
	}

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	time_now := time.Now()

	user := auth_entity.NewUser(
		cuid.New(),
		username,
		string(hashedPassword),
		time_now,
		time_now,
	)

	// ユーザーをデータベースに保存
	err = s.UserRepo.Create(&user)
	if err != nil {
		return "", err
	}

	// JWTを生成して返す
	token, err := GenerateJWT(user.ID())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) LoginUser(username, password string) (string, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid username or password")
	}

	// パスワードの照合
	err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// JWT トークンを生成
	token, err := GenerateJWT(user.ID())
	if err != nil {
		return "", err
	}

	return token, nil
}
