package vo

import (
	"fmt"
	"todoapps/domain_todocore/model/errorenum"
)

type TodoID string

func NewTodoID(random6Char string) (TodoID, error) {

	if len(random6Char) != 6 {
		return "", errorenum.Random6CharLengthNotSatisfied
	}

	obj := TodoID(fmt.Sprintf("TODO-%s", random6Char))

	return obj, nil
}

func (r TodoID) Validate() error {
	return nil
}

func (r TodoID) String() string {
	return string(r)
}
