package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	database *gorm.DB
	logger   *Logger
)

func Init() error {
	var initializeDatabaseError error
	database, initializeDatabaseError = InitializeSQLite()
	if initializeDatabaseError != nil {
		return fmt.Errorf("error initializing swlite %w", initializeDatabaseError)
	}

	return nil
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}
