// presentation/router.go
package router

import (
	auth_services "myproject/application/auth"
	"myproject/application/services"
	auth_handler "myproject/presentation/auth"
	"myproject/presentation/handler"

	"github.com/labstack/echo/v4"
)

// main.goからここに飛ぶ
func NewRouter(engine *echo.Echo, service *services.Service, auth *auth_services.AuthService) {

	//認証機能のハンドラー設定
	authHandler := auth_handler.NewAuthHandler(auth)
	h := &auth_handler.AuthHandler{}

	//ハンドラのインスタンスを作成
	todoHandler := handler.NewToDoHandler(service)
	staffHandler := handler.NewStaffHandler(service)
	todostaffHandler := handler.NewToDo_StaffHandler(service)

	//todoハンドラ
	engine.GET("/todos", h.JWTMiddleware(todoHandler.GetToDo))
	engine.GET("/todos/:id", h.JWTMiddleware(todoHandler.GetByID))
	engine.POST("/todos", h.JWTMiddleware(todoHandler.CreateToDo))
	engine.PATCH("/todos/:id", h.JWTMiddleware(todoHandler.UpdateToDo))
	engine.GET("/todos/:id/is_completed", h.JWTMiddleware(todoHandler.CompletedToDo))
	engine.POST("/todos/:id/duplicate", h.JWTMiddleware(todoHandler.DuplicateToDo))
	engine.DELETE("/todos/:id", h.JWTMiddleware(todoHandler.DeleteToDo))

	// staffハンドラ
	engine.GET("/staffs", h.JWTMiddleware(staffHandler.GetStaff))
	engine.GET("/staffs/:id", h.JWTMiddleware(staffHandler.GetByStaff_ID))
	engine.POST("/staffs", h.JWTMiddleware(staffHandler.CreateStaff))
	engine.PATCH("/staffs/:id", h.JWTMiddleware(staffHandler.UpdateStaff))
	engine.DELETE("/staffs/:id", h.JWTMiddleware(staffHandler.DeleteStaff))

	// todo_staffハンドラ
	engine.GET("/todo_staff", h.JWTMiddleware(todostaffHandler.GetAssign))
	engine.GET("/todo_staff/:id", h.JWTMiddleware(todostaffHandler.GetAssignByID))
	engine.POST("/todo_staff", h.JWTMiddleware(todostaffHandler.Assigntodo))
	engine.PATCH("/todo_staff/:id", h.JWTMiddleware(todostaffHandler.UpdateAssign))
	engine.DELETE("/todo_staff/:id", h.JWTMiddleware(todostaffHandler.DeleteAssign))

	// 認証機能
	engine.POST("/register", authHandler.Register)
	engine.POST("/login", authHandler.Login)
}
