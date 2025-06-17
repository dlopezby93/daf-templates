//go:build wireinject
// +build wireinject

package di

import (
	"{{module_name}}/internal/api/controllers"
	"{{module_name}}/internal/application/config"
	"{{module_name}}/internal/application/services"
	"{{module_name}}/internal/infrastructure/clients"
	"{{module_name}}/internal/infrastructure/repositories"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	TransactionTypeController *controllers.TransactionTypeController
}

func InitializeApp() (*App, error) {
	wire.Build(
		config.NewApplicationConfig,
		services.NewTransactionTypeApplicationService,
		repositories.NewTransactionTypeRepository,
		ProvideMongoDBClient,
		controllers.NewTransactionTypeController,
		wire.Struct(new(App), "TransactionTypeController"),
	)

	return &App{}, nil
}

func ProvideMongoDBClient(config *config.ApplicationConfig) *mongo.Database {
	mongodb := clients.NewMongoClient(config.MongoURI).GetClient().Database(config.Database)

	return mongodb
}
