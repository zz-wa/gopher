package main

import (
	"Q_two_goose/database"
	"Q_two_goose/handle"

	"go.uber.org/zap"
)

func main() {

	db, err := database.Connect("identifier.sqlite")

	if err != nil {
		zap.S().Fatal("database connect failed", zap.Error(err))
	}

	server := handle.NewServer(db)
	if err := server.RUN(); err != nil {
		zap.S().Fatal("server run failed", zap.Error(err))
	}

}
