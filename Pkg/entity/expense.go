package entity

import "time"

// Expense representa los movimientos
type Expense struct {
	UserID      int
	Category    string
	Name        string
	Description string
	Amount      float32
	Date        time.Time
}
