package runtodocreate

import (
	"context"
	"time"

	"todoapps/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	Message string
	Now     time.Time
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}
