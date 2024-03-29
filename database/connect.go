package database

import (
	"log/slog"
	"os"

	"github.com/Nohty/blog/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const USER_ID = 1

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		slog.Error("Failed to connect to database")
		os.Exit(1)
	}

	slog.Info("Connected to database")
	DB.AutoMigrate(&model.User{}, &model.BlogPost{})
	slog.Info("Database migrated")
}
