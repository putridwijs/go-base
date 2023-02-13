package greeting

import (
	"context"
	"go-base/internal/pkg/greeting/dto"
)

// UseCase interface of greeting package
type UseCase interface {
	Greet(ctx context.Context, request dto.GreetingRequest) (dto.GreetingResponse, error)
}
