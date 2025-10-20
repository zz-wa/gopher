package handle

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	database *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{database: db}
}

func (s *Server) RUN() error {
	router := echo.New()
	router.HideBanner = true
	router.HidePort = true

	router.Use(middleware.Recover())

	router.POST("/urls", s.Insert)

	router.GET("/:path", s.MapHandler)
	return router.Start(":8080")

}

//这里要进行重定向操作

func (s *Server) MapHandler(c echo.Context) error {
	/*path := c.Request().URL.Path[1:]
	 */

	path := c.Param("path")
	if path == "" {
		return c.String(http.StatusOK, "Hello, World!")
	}

	var originalURL string
	if err := s.database.QueryRow("SELECT OriginalURL FROM url_Shorter WHERE ShortURL = ?", path).Scan(&originalURL); err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.Redirect(http.StatusMovedPermanently, originalURL)
}

type InsertReq struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type InsertRes struct {
	ShortURL string `json:"short_url"`
}

func (S *Server) Insert(c echo.Context) error {
	var request InsertReq
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("JSON bind failed: %v", err),
		})
	}
	if request.OriginalURL == "" || request.ShortURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_url and original_url required"})
	}

	result, err := S.database.Exec("INSERT INTO url_Shorter (OriginalURL,ShortURL) VALUES (?,?) ", request.OriginalURL, request.ShortURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Database insert failed: %v", err),
		})
	}

	_, err = result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to get last insert id: %v", err),
		})
	}
	return c.JSON(http.StatusCreated, InsertRes{ShortURL: request.ShortURL})

}
