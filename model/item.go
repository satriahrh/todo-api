package model

import "time"

type Item struct {
	ID          int
	Name        string
	Description string
	IsDone      bool
	CreatedAt   time.Time
}
