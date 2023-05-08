package domain

import "time"

type FinancialPlanner struct {
	Id            int
	User_id       float64
	Planner_name  string
	Planner_email string
	Created_at    time.Time
	Updated_at    time.Time
}
