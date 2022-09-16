package restapi

import (
	"github.com/gin-gonic/gin"

	"todoapps/domain_todocore/usecase/getalltodo"
	"todoapps/domain_todocore/usecase/runtodocheck"
	"todoapps/domain_todocore/usecase/runtodocreate"
	"todoapps/shared/infrastructure/config"
	"todoapps/shared/infrastructure/logger"
)

type Controller struct {
	Router              gin.IRouter
	Config              *config.Config
	Log                 logger.Logger
	GetAllTodoInport    getalltodo.Inport
	RunTodoCheckInport  runtodocheck.Inport
	RunTodoCreateInport runtodocreate.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/todo", r.authorized(), r.getAllTodoHandler())
	r.Router.PUT("/todo/:todo_id", r.authorized(), r.runTodoCheckHandler())
	r.Router.POST("/todo", r.authorized(), r.runTodoCreateHandler())
}
