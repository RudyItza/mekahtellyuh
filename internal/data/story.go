package data

import "time"

type Stories struct {
	ID        int
	Title     string
	Content   string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
