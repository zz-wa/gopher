package model

type ShorterURL struct {
	Number      int    `gorm:"primary_key" json:"number"`
	OriginalURL string `gorm:"type:text" json:"originalURL"`
	ShorterURL  string `gorm:"type:text" json:"shorterURL"`
}
