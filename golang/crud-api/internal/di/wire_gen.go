// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"{{module_name}}/internal/api/controllers"
	"{{module_name}}/internal/application/config"
	"{{module_name}}/internal/application/services"
	"{{module_name}}/internal/infrastructure/clients"
	"{{module_name}}/internal/infrastructure/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitializeApp() (*App, error) {
	applicationConfig := config.NewApplicationConfig()
	database := ProvideMongoDbClient(applicationConfig)
	iTransactionTypeRepository := repositories.NewTransactionTypeRepository(database)
	iTransactionTypeApplicationService := services.NewTransactionTypeApplicationService(iTransactionTypeRepository)
	transactionTypeController := controllers.NewTransactionTypeController(iTransactionTypeApplicationService)
	app := &App{
		TransactionTypeController: transactionTypeController,
	}
	
	return app, nil
}

// wire.go:

type App struct {
	TransactionTypeController *controllers.TransactionTypeController
}

func ProvideMongoDbClient(config2 *config.ApplicationConfig) *mongo.Database {
	mongodb := clients.NewMongoClient(config2.MongoURI).GetClient().Database(config2.Database)

	return mongodb
}