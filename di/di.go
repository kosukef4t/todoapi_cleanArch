package di

import (
	"database/sql"
	auth_services "myproject/application/auth"
	integration_services "myproject/application/integration"
	"myproject/application/services"
	"myproject/infrastructure/auth"
	repository "myproject/infrastructure/gateway"
)

// InitializeServices は手動で依存関係を初期化します
func InitializeServices(db *sql.DB) *integration_services.Services {
	// リポジトリの初期化
	isToDoRepository := repository.NewToDoRepository(db)
	isStaffRepository := repository.NewStaffRepository(db)
	isToDoStaffRepository := repository.NewToDoStaffRepository(db)
	isUserRepository := auth.NewUserRepository(db)

	// サービスの初期化
	service := services.NewService(isToDoRepository, isStaffRepository, isToDoStaffRepository)
	authService := auth_services.NewAuthService(isUserRepository)

	// 依存関係を統合して返す
	return &integration_services.Services{
		Service:     service,
		AuthService: authService,
	}
}
