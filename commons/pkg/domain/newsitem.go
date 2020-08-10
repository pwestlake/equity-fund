package domain

import (
	"time"
)

// NewsItem ...
// Domain struct for a NewsItem
type NewsItem struct {
	ID          string    `json:"id"`
	CatalogRef  string    `json:"catalogref"`
	CompanyCode string    `json:"companycode"`
	CompanyName string    `json:"companyname"`
	DateTime    time.Time `json:"datetime"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Sentiment   float32   `json:"sentiment"`
}
