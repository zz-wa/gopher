package model

import "gorm.io/gorm"

type URLShorter struct {
	gorm.Model
	OriginalURL string `gorm:"not null" json:"original_url"`
	ShortURL    string `gorm:"not null" json:"short_url"`
}
