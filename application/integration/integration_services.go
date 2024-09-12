package integration_services

// Services struct to bundle ToDoService and AuthService
import (
	auth_services "myproject/application/auth"
	"myproject/application/services"
)

type Services struct {
	Service     *services.Service
	AuthService *auth_services.AuthService
}
