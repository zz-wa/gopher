package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)


func Connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		zap.S().Fatal("failed to connect database", zap.Error(err))
	}

	if err := db.Ping(); err != nil {
		zap.S().Fatal("failed to ping database", zap.Error(err))
	}
	goose.SetDialect("sqlite3")

	if err := goose.Up(db, "./migrations"); err != nil {
		zap.S().Fatal("failed to apply migrations", zap.Error(err))
	}

	zap.L().Info("Connected to database", zap.String("path", path))
	return db, nil
}
