package domain

import "time"

type Category struct {
	Id         int
	Name       string
	Type       Flows
	Created_at time.Time
	Updated_at time.Time
}
