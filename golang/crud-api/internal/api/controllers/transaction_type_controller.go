package controllers

import (
	"errors"
	"net/http"

	localErrors "{{module_name}}/internal/application/errors"
	"{{module_name}}/internal/application/services"

	"github.com/gin-gonic/gin"
)

type TransactionTypeController struct {
	appService services.ITransactionTypeApplicationService
}

func NewTransactionTypeController(appService services.ITransactionTypeApplicationService) *TransactionTypeController {
	return &TransactionTypeController{appService: appService}
}

func (controller *TransactionTypeController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/catalog")

	api.GET("/transaction-types", controller.GetAllTransactionTypes)
	api.GET("/transaction-types/:id", controller.GetTransactionTypeById)
}

func (controller *TransactionTypeController) GetAllTransactionTypes(ctx *gin.Context) {
	response, err := controller.appService.GetAll()

	if err != nil {
		ctx.JSON(resolveStatus(err), gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *TransactionTypeController) GetTransactionTypeById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := controller.appService.GetById(id)

	if err != nil {
		ctx.JSON(resolveStatus(err), gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func resolveStatus(err error) int {
	switch {
	case errors.Is(err, localErrors.ErrInvalidInput):
		return http.StatusBadRequest
	case errors.Is(err, localErrors.ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
