package application

import (
	"todoapps/domain_todocore/controller/restapi"
	"todoapps/domain_todocore/gateway/sqlitedb"
	"todoapps/domain_todocore/usecase/getalltodo"
	"todoapps/domain_todocore/usecase/runtodocheck"
	"todoapps/domain_todocore/usecase/runtodocreate"
	"todoapps/shared/driver"
	"todoapps/shared/infrastructure/config"
	"todoapps/shared/infrastructure/logger"
	"todoapps/shared/infrastructure/server"
	"todoapps/shared/infrastructure/util"
)

type apptodo struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c apptodo) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewApptodo() func() driver.RegistryContract {

	const appName = "apptodo"

	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData(appName, appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

		//datasource := inmemory.NewGateway(log, appData, cfg)
		//datasource := mongodb.NewGateway(log, appData, cfg)
		datasource := sqlitedb.NewGateway(log, appData, cfg)

		return &apptodo{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                 log,
				Config:              cfg,
				Router:              httpHandler.Router,
				GetAllTodoInport:    getalltodo.NewUsecase(datasource),
				RunTodoCheckInport:  runtodocheck.NewUsecase(datasource),
				RunTodoCreateInport: runtodocreate.NewUsecase(datasource),
			},
		}

	}
}
