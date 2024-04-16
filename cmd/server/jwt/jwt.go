package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	// Секретный ключ для подписи токена
	secretKey = []byte("secret_key_here")
)

// Структура для хранения информации о пользователе
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Генерация JWT-токена
func generateToken(user User) (string, error) {
	// Устанавливаем срок действия токена
	expirationTime := time.Now().Add(24 * time.Hour)

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	})

	// Подписываем токен
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Проверка валидности JWT-токена
func validateToken(tokenString string) (*jwt.Token, error) {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
