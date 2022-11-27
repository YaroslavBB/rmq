package global

import "errors"

var (
	// ErrNoData данные не найдены
	ErrNoData = errors.New("данные не найдены")
	// ErrIncorrectParam неверные параметры
	ErrIncorrectParam = errors.New("неверные параметры")
	// ErrInternalError сервер временно недоступен
	ErrInternalError = errors.New("сервер временно недоступен")
)
