package domain

import (
	"errors"
	"time"
)

type Flows int64

const (
	Income Flows = iota
	Outcome
)

func ToFlowData(flow string) (Flows, error) {
	if flow == "income" {
		return Income, nil
	} else if flow == "outcome" {
		return Outcome, nil
	}
	return 440, errors.New("salah input tipe data")
}

func ToFlowString(flow Flows) (string, error) {
	if flow == Income {
		return "income", nil
	} else if flow == Outcome {
		return "outcome", nil
	}
	return "", errors.New("harus input data flows")
}

type Transaction struct {
	Id          int
	User_id     int
	Category_id int
	Amount      float32
	Description string
	Date        string
	Type        Flows
	Created_at  time.Time
	Updated_at  time.Time
}
