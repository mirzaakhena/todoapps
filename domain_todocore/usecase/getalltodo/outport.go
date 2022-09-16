package getalltodo

import (
	"todoapps/domain_todocore/model/repository"
)

// Outport of usecase
type Outport interface {
	repository.FindAllTodoRepo
}
