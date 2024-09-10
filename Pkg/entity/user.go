package entity

import "time"

// Siempre comentar todo
type User struct {
	ID        int
	Name      string
	LastName  string
	Email     string
	Password  string
	CreatedAt *time.Time
	DeletedAt *time.Time
}
