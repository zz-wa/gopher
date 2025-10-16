package database

import (
	"Q_two/database/model"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	database *sql.DB
}

type InsertReq struct {
	number      int    `json:"number"`
	ShorterURL  string `json:"shorterURL"`
	originalURL string `json:"original_url"`
}
type InsertRes struct {
	shorterURL string
}

// 插入的，但是在那里用捏
func (S *Server) Insert(c echo.Context) error {
	var request InsertReq

	if err := c.Bind(&request); err != nil {
		return err
	}
	var shorter model.ShorterURL
	if err := S.database.QueryRow("INSERT INTO urlShorter (originalURL, shorterURL) VALUES (? ?) RETURNING number", request.originalURL, request.originalURL).
		Scan(&shorter.Number); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, InsertRes{
		shorterURL: shorter.ShorterURL,
	})
}
