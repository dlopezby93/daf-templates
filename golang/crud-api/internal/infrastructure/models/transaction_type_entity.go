package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionTypeEntity struct {
	ID                        primitive.ObjectID    `bson:"_id"`
	Name                      string                `bson:"name"`
	Deleted                   bool                  `bson:"deleted"`
	DefinitionType            string                `bson:"definitionType"`
	CreatedAt                 time.Time             `bson:"createdAt"`
	ReversalTransactionType   string                `bson:"reversalTransactionType"`
	CommissionTransactionType string                `bson:"commissionTransactionType"`
	CommissionType            string                `bson:"commissionType"`
	FiservActionCode          FiservActionCode      `bson:"fiservActionCode"`
	StatementCategory         StatementCategory     `bson:"statementCategory"`
	ApplicationActionCategory string                `bson:"applicationActionCategory"`
	AccountLimitsValidation   string                `bson:"accountLimitsValidation"`
	PostExecutionEvents       []PostExecutionEvents `bson:"postExecutionEvents"`
	Disputable                string                `bson:"disputable"`
}
