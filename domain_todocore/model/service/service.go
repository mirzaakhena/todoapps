package service

import "context"

type GenerateRandomIDService interface {
	GenerateRandomChar(ctx context.Context) string
}
