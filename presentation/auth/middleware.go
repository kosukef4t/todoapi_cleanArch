package auth_handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserContextKey struct{}

func (h *AuthHandler) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Authorizationヘッダーからトークンを取得
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Authorization header required")
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// トークンのパースと検証
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("your_secret_key"), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		// トークンのクレームが有効であれば、コンテキストにユーザーIDを格納
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// EchoのコンテキストにユーザーIDを保存
			c.Set("user_id", claims["user_id"])
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}
	}
}
