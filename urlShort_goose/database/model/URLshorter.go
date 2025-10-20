package model

type URLShorter struct {
	ID          int    `param:"id" json:"id"`
	OriginalURL string ` param:"original_url"  json:"original_url"`
	ShortURL    string `param:"short_url"  json:"short_url"`
}
