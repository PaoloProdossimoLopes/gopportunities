package config

import (
	"os"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	const dbPath = "./db/main.db"
	logger := GetLogger("sqlite")

	_, statError := os.Stat(dbPath)
	if os.IsNotExist(statError) {
		logger.Info("database file not found, creating ...")

		crateDatabaseFolderError := os.MkdirAll("./db", os.ModePerm)
		if crateDatabaseFolderError != nil {
			logger.Errorf("Fail on create database folder: %v", crateDatabaseFolderError)
			return nil, crateDatabaseFolderError
		}

		databaseFile, createDatabaseFileError := os.Create(dbPath)
		if createDatabaseFileError != nil {
			logger.Errorf("Fail on create database file: %v", crateDatabaseFolderError)
			return nil, createDatabaseFileError
		}

		databaseFile.Close()
	}

	database, openDataBaseError := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if openDataBaseError != nil {
		logger.Errorf("opening database error: %v", openDataBaseError)
		return nil, openDataBaseError
	}

	migrateDataBaseError := database.AutoMigrate(&schemas.Opening{})
	if migrateDataBaseError != nil {
		logger.Errorf("Fail on auto migration database: %v", migrateDataBaseError)
		return nil, migrateDataBaseError
	}

	return database, nil
}
