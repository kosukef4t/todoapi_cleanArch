package auth_handler

import (
	auth_services "myproject/application/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService *auth_services.AuthService
}

func NewAuthHandler(userservice *auth_services.AuthService) *AuthHandler {
	return &AuthHandler{userService: userservice}
}

// dtoの構造体
type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ユーザー登録のHTTPハンドラー
func (h *AuthHandler) Register(c echo.Context) error {
	var req Request
	// リクエストボディをデコードして `req` にバインド
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	// ユーザー登録
	token, err := h.userService.RegisterUser(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// JWTトークンを返す
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req Request
	// リクエストボディをデコードして `req` にバインド
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// ユーザーログイン
	token, err := h.userService.LoginUser(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	// JWTトークンを返す
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
