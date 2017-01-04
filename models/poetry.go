package models

import (
	"time"
)

type Poetry struct {
	Id        int
	PoetId    int
	Content   string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
