package mappers

import (
	"{{module_name}}/internal/application/dtos"
	"{{module_name}}/internal/infrastructure/models"
)

func TransactionTypeEntityToDto(entity models.TransactionTypeEntity) dtos.TransactionTypeDto {
	return dtos.TransactionTypeDto{
		ID:                        entity.ID.Hex(),
		Name:                      entity.Name,
		Deleted:                   entity.Deleted,
		DefinitionType:            entity.DefinitionType,
		CreatedAt:                 entity.CreatedAt,
		ReversalTransactionType:   entity.ReversalTransactionType,
		CommissionTransactionType: entity.CommissionTransactionType,
		CommissionType:            entity.CommissionType,
		FiservActionCode:          dtos.FiservActionCode(entity.FiservActionCode),
		StatementCategory:         dtos.StatementCategory(entity.StatementCategory),
		ApplicationActionCategory: entity.ApplicationActionCategory,
		AccountLimitsValidation:   entity.AccountLimitsValidation,
		PostExecutionEvents:       PostExecutionEventsToDto(entity.PostExecutionEvents),
		Disputable:                entity.Disputable,
	}
}

func PostExecutionEventsToDto(entity []models.PostExecutionEvents) []dtos.PostExecutionEvents {
	var postExecutionEventsDtos []dtos.PostExecutionEvents
	for _, postExecutionEvent := range entity {
		postExecutionEventsDtos = append(postExecutionEventsDtos, PostExecutionEventToDto(postExecutionEvent))
	}

	return postExecutionEventsDtos
}

func PostExecutionEventToDto(entity models.PostExecutionEvents) dtos.PostExecutionEvents {
	return dtos.PostExecutionEvents{
		Event: entity.Event,
	}
}
