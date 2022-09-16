package runtodocheck

import "todoapps/domain_todocore/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindOneTodoRepo
	repository.SaveTodoRepo
}
