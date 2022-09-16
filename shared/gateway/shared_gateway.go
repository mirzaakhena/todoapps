package gateway

import (
	"context"
	"todoapps/shared/util"
)

type SharedGateway struct {
}

func (r *SharedGateway) GenerateRandomChar(ctx context.Context) string {
	return util.GenerateID(6)
}
