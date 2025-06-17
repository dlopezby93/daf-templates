package models

type StatementCategory struct {
	Name                   string `json:"name"`
	AccountBalanceModifier int    `json:"accountBalanceModifier"`
	CategoryId             int    `json:"categoryId"`
}
