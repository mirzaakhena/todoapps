package sqlitedb

import (
	"context"
	"todoapps/domain_todocore/model/entity"
	"todoapps/shared/driver"
	gateway2 "todoapps/shared/gateway"
	"todoapps/shared/infrastructure/config"
	"todoapps/shared/infrastructure/database"
	"todoapps/shared/infrastructure/logger"
)

type gateway struct {
	*gateway2.SharedGateway
	*database.GormWithoutTransactionImpl

	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	gwt := database.NewGormWithoutTransactionImpl(database.NewSQLiteDefault())

	err := gwt.DB.AutoMigrate(entity.Todo{})
	if err != nil {
		panic(err.Error())
	}

	return &gateway{
		SharedGateway:              &gateway2.SharedGateway{},
		GormWithoutTransactionImpl: gwt,
		log:                        log,
		appData:                    appData,
		config:                     cfg,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	var count int64
	err := r.DB.Model(&entity.Todo{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	results := make([]*entity.Todo, 0)
	err = r.DB.Find(&results).Error
	if err != nil {
		return nil, 0, err
	}

	return results, count, nil
}

func (r *gateway) FindOneTodo(ctx context.Context, todoID string) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	var result entity.Todo
	err := r.DB.First(&result, "id = ?", todoID).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	err := r.DB.Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}
