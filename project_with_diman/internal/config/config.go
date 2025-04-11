package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	EmptyError    = errors.New("data is empty")
	EmptyKeyError = errors.New("key is empty")
	secretKey     string
)

func Config() (string, error) {
	user := os.Getenv("user")
	password := os.Getenv("password")
	if user == "" || password == "" {
		return "", fmt.Errorf("%w", EmptyError)
	}
	return fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=postgres sslmode=disable", user, password), nil
}

// CreateKey - функция, получает секретный ключ для подписи jwt
func CreateKey() (string, error) {
	secretKey = os.Getenv("secret_key")
	if secretKey == "" {
		return "", fmt.Errorf("key is empty %w", EmptyKeyError)
	}
	return secretKey, nil
}

func GetKey() (string, error) {
	if secretKey == "" {
		return "", fmt.Errorf("key is empty %w", EmptyKeyError)
	}
	return secretKey, nil
}
