package domain

import "time"

type Projections int64

const (
	Monthly Projections = iota
	Yearly
)

type FinancialProjections struct {
	Id         int
	User_id    int
	Type       Projections
	Start_date time.Time
	End_date   time.Time
	Amount     float64
	Created_at time.Time
	Updated_at time.Time
}
