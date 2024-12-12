package models

type FiservActionCode struct {
	Code                  string `json:"code"`
	Api                   string `json:"api"`
	TransactionIdentifier string `json:"transactionIdentifier"`
	OperationType         string `json:"operationType"`
}
