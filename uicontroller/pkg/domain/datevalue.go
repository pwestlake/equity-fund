package domain

import (
	"time"
)

// DateValue ..
// DTO
type DateValue struct {
	Date  time.Time `json:"date"`
	Value float32   `json:"value"`
}
