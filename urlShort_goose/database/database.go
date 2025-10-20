package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

// 连接数据库
// 现在没有迁移的相关代码

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

/*
func RunMigrations(db *sql.DB) error {
	if err := goose.RunMigrations(db, "sqlite", "./migrations", false); err != nil {
		return fmt.Errorf("migrations failed: %w", err)
	}
	return nil
}
*/
