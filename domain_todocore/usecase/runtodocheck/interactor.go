package runtodocheck

import (
	"context"
)

//go:generate mockery --name Outport -output mocks/

type runTodoCheckInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runTodoCheckInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runTodoCheckInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := r.outport.FindOneTodo(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}

	err = todoObj.CheckDone()
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
