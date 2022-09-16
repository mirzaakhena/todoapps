package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"todoapps/domain_todocore/model/entity"
	"todoapps/shared/driver"
	gateway2 "todoapps/shared/gateway"
	"todoapps/shared/infrastructure/config"
	"todoapps/shared/infrastructure/database"
	"todoapps/shared/infrastructure/logger"
)

type gateway struct {
	*gateway2.SharedGateway
	*database.MongoWithTransaction

	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	databaseName := "tododb"
	uri := fmt.Sprintf("mongodb://localhost:27017/%s?readPreference=primary&ssl=false", databaseName)
	mwt := database.NewMongoWithTransaction(database.NewMongoDefault(uri), databaseName)

	mwt.PrepareCollection(
		entity.Todo{},
	)

	return &gateway{
		SharedGateway:        &gateway2.SharedGateway{},
		MongoWithTransaction: mwt,
		log:                  log,
		appData:              appData,
		config:               cfg,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	filter := bson.M{}
	results := make([]*entity.Todo, 0)
	count, err := r.GetAll(ctx, page, size, filter, &results)
	if err != nil {
		return nil, 0, err
	}

	return results, count, nil
}

func (r *gateway) FindOneTodo(ctx context.Context, todoID string) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	var result entity.Todo
	err := r.GetOne(ctx, todoID, &result)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	return &result, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	_, err := r.SaveOrUpdate(ctx, string(obj.ID), obj)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return err
	}

	return nil
}
