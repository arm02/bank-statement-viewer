package model

import "time"

type TxType string

const (
	Credit TxType = "CREDIT"
	Debit  TxType = "DEBIT"
)

type Status string

const (
	Success Status = "SUCCESS"
	Failed  Status = "FAILED"
	Pending Status = "PENDING"
)

type Transaction struct {
	Timestamp   time.Time `json:"timestamp"`
	Name        string    `json:"name"`
	Type        TxType    `json:"type"`
	Amount      int64     `json:"amount"`
	Status      Status    `json:"status"`
	Description string    `json:"description"`
}
