package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("fake erros")
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}
