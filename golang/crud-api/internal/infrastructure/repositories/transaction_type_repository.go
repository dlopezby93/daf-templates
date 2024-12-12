package repositories

import (
	"context"
	"time"

	"{{module_name}}/internal/application/errors"
	"{{module_name}}/internal/infrastructure/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TransactionTypesCollection = "transactionTypes"
)

type ITransactionTypeRepository interface {
	GetAll() ([]models.TransactionTypeEntity, error)
	GetById(id string) (models.TransactionTypeEntity, error)
}

type MongoTransactionTypeRepository struct {
	collection *mongo.Collection
}

var _ ITransactionTypeRepository = &MongoTransactionTypeRepository{}

func NewTransactionTypeRepository(db *mongo.Database) ITransactionTypeRepository {
	return &MongoTransactionTypeRepository{collection: db.Collection(TransactionTypesCollection)}
}

func (repository *MongoTransactionTypeRepository) GetAll() ([]models.TransactionTypeEntity, error) {
	var transactionTypes []models.TransactionTypeEntity

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repository.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.ErrDB
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var transactionType models.TransactionTypeEntity

		if err := cursor.Decode(&transactionType); err != nil {
			return nil, errors.WrapError(errors.ErrMappingData, err.Error())
		}
		transactionTypes = append(transactionTypes, transactionType)
	}

	return transactionTypes, nil
}

func (repository *MongoTransactionTypeRepository) GetById(id string) (models.TransactionTypeEntity, error) {
	var transactionType models.TransactionTypeEntity

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return transactionType, errors.WrapError(errors.ErrInvalidInput, err.Error())
	}

	err = repository.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&transactionType)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errors.WrapError(errors.ErrNotFound, err.Error())
		} else {
			err = errors.ErrDB
		}
		return transactionType, err
	}

	return transactionType, nil
}
