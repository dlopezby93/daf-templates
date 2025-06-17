package dtos

import (
	"time"
)

type TransactionTypeDto struct {
	ID                        string                `json:"id"`
	Name                      string                `json:"name"`
	Deleted                   bool                  `json:"deleted"`
	DefinitionType            string                `json:"definitionType"`
	CreatedAt                 time.Time             `json:"createdAt"`
	ReversalTransactionType   string                `json:"reversalTransactionType"`
	CommissionTransactionType string                `json:"commissionTransactionType"`
	CommissionType            string                `json:"commissionType"`
	FiservActionCode          FiservActionCode      `json:"fiservActionCode"`
	StatementCategory         StatementCategory     `json:"statementCategory"`
	ApplicationActionCategory string                `json:"applicationActionCategory"`
	AccountLimitsValidation   string                `json:"accountLimitsValidation"`
	PostExecutionEvents       []PostExecutionEvents `json:"postExecutionEvents"`
	Disputable                string                `json:"disputable"`
}
