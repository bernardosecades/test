package entity

import "time"

type Product struct {
	ID        string
	Available bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
