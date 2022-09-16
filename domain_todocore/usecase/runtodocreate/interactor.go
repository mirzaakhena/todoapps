package runtodocreate

import (
	"context"
	"todoapps/domain_todocore/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type runTodoCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runTodoCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runTodoCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	randomID := r.outport.GenerateRandomChar(ctx)

	todoObj, err := entity.NewTodo(entity.TodoRequest{
		Random6Char: randomID,
		Message:     req.Message,
		Now:         req.Now,
	})
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
