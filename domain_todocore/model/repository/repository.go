package repository

import (
	"context"
	"todoapps/domain_todocore/model/entity"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FindOneTodoRepo interface {
	FindOneTodo(ctx context.Context, todoID string) (*entity.Todo, error)
}

type FindAllTodoRepo interface {
	FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error)
}

type FindAllTodoBlbalbaRepo interface {
	FindAllTodoBlbalba(ctx context.Context, page, size int64, someID string) ([]*entity.Todo, int64, error)
}
