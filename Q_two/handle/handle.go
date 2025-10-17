package handle

import (
	"Q_two/database/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server struct {
	database *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{database: db}
}

func (s *Server) RUN() error {
	router := echo.New()
	router.HideBanner = true
	router.HidePort = true

	router.Use(middleware.Recover())

	router.POST("/urls", s.Insert)

	router.GET("/*", s.MapHandler)
	return router.Start(":8080")

}

//这里要进行重定向操作

func (s *Server) MapHandler(c echo.Context) error {
	path := c.Request().URL.Path[1:]

	if path == "" || path == "/" {
		return c.String(http.StatusOK, "Hello, World!")
	}

	var url model.URLShorter
	if err := s.database.Where("short_url = ?", path).First(&url).Error; err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.Redirect(http.StatusFound, url.OriginalURL)

}

type InsertReq struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type InsertRes struct {
	ShortURL string `json:"short_url"`
}

// 插入的，但是在那里用捏

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

	//插入语句

	if err := S.database.Create(&model.URLShorter{
		OriginalURL: request.OriginalURL,
		ShortURL:    request.ShortURL,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, InsertRes{ShortURL: request.ShortURL})

}
