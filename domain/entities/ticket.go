package entities

import "time"

type Ticket struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
