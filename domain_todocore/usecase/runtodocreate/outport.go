package runtodocreate

import (
	"todoapps/domain_todocore/model/repository"
	"todoapps/domain_todocore/model/service"
)

// Outport of usecase
type Outport interface {
	repository.SaveTodoRepo
	service.GenerateRandomIDService
}
