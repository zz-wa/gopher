package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func Connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	zap.L().Info("connected to database")
	return db, nil
}

func GetURL(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query("SELECT originalURL, shorterURL   FROM urlShorter")
	if err != nil {
		zap.L().Error("error getting url", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	urls := make(map[string]string)

	for rows.Next() {
		var shorterURL, originalURL string
		err := rows.Scan(&shorterURL, &originalURL)
		if err != nil {
			zap.L().Error("error getting url", zap.Error(err))
		}
		urls[originalURL] = shorterURL
	}
	return urls, nil
}
