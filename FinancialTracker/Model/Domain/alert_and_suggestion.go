package domain

import "time"

type Enum int64

const (
	Alert Enum = iota
	Suggestion
)

type FinancialAlertAndSuggestion struct {
	Id         int
	User_id    int
	Type       Enum
	Message    string
	Created_at time.Time
	Updated_at time.Time
}
