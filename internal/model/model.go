package model

import (
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("not found identifier in database")
	ErrIdentifierIsExist = errors.New("identifier already exist")
)

type Shortening struct {
	Identifier  string    `json:"identifier"`
	OriginalUrl string    `json:"original_url"`
	Visits      string    `json:"visits"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
