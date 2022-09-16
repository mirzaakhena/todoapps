package entity

import (
	"time"
	"todoapps/domain_todocore/model/errorenum"
	"todoapps/domain_todocore/model/vo"
)

type Todo struct {
	ID      vo.TodoID `json:"id" bson:"_id"`
	Message string    `json:"message" bson:"message"`
	Date    time.Time `json:"date" bson:"date"`
	Done    bool      `json:"done" bson:"done"`
}

type TodoRequest struct {
	Random6Char string
	Message     string
	Now         time.Time
}

func NewTodo(req TodoRequest) (*Todo, error) {

	// validate the message
	if req.Message == "" {
		return nil, errorenum.MessageMustNotEmpty
	}

	// create the id
	id, err := vo.NewTodoID(req.Random6Char)
	if err != nil {
		return nil, err
	}

	// prepare the object
	var obj Todo

	// assign the value
	obj.ID = id
	obj.Message = req.Message
	obj.Date = req.Now

	return &obj, nil
}

func (r *Todo) CheckDone() error {
	if r.Done {
		return errorenum.CannotCheckTheDoneTodoItem
	}
	r.Done = true
	return nil
}
