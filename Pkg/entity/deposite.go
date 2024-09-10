package entity

import "time"

//Deposite para los ingresos
type Deposite struct {
	UserID      int
	Category    string
	Name        string
	Description string
	Amount      float32
	Date        time.Time
}
