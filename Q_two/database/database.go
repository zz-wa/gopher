package database

import (
	"Q_two/database/model"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//连接数据库并且自动创建一个数据库，利用结构体映射

func Connect(path string) (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	//db, err := database.Connect("identifier.sqlite")
	if err != nil {
		zap.S().Fatal("failed to connect to database", zap.Error(err))
	}

	if err = db.AutoMigrate(&model.URLShorter{}); err != nil {
		zap.S().Fatal("failed to migrate table", zap.Error(err))
	}

	zap.S().Info("successfully connected to database")

	return db
}
