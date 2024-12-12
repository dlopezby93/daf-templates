package services

import (
	"{{module_name}}/internal/application/dtos"
	"{{module_name}}/internal/application/mappers"
	"{{module_name}}/internal/infrastructure/models"
	"{{module_name}}/internal/infrastructure/repositories"

	spinmappers "github.com/fintechdigitalventure/go-spin-common/mappers"
)

type ITransactionTypeApplicationService interface {
	GetAll() ([]dtos.TransactionTypeDto, error)
	GetById(id string) (dtos.TransactionTypeDto, error)
}

type TransactionTypeApplicationService struct {
	mapper                    spinmappers.MapFunc[models.TransactionTypeEntity, dtos.TransactionTypeDto]
	transactionTypeRepository repositories.ITransactionTypeRepository
}

var _ ITransactionTypeApplicationService = &TransactionTypeApplicationService{}

func NewTransactionTypeApplicationService(repository repositories.ITransactionTypeRepository) ITransactionTypeApplicationService {
	return &TransactionTypeApplicationService{mapper: mappers.TransactionTypeEntityToDto, transactionTypeRepository: repository}
}

func (appService *TransactionTypeApplicationService) GetAll() ([]dtos.TransactionTypeDto, error) {
	var response []dtos.TransactionTypeDto
	transactionTypeEntities, err := appService.transactionTypeRepository.GetAll()

	if err == nil {
		response, err = appService.mapper.MapEachErr(transactionTypeEntities, err)
		if err == nil {
			return response, nil
		}
	}

	return response, err
}

func (appService *TransactionTypeApplicationService) GetById(id string) (dtos.TransactionTypeDto, error) {
	var response dtos.TransactionTypeDto
	transactionTypeEntity, err := appService.transactionTypeRepository.GetById(id)

	if err == nil {
		response, err = appService.mapper.MapErr(transactionTypeEntity, err)
		if err == nil {
			return response, nil
		}
	}
	return response, err
}
