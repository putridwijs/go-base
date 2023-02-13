package endpoint

import (
	"go-base/internal/config"
	"go-base/internal/pkg/greeting"
	"go-base/internal/pkg/greeting/usecase"
)

type Endpoint struct {
	cfg     config.Config
	useCase greeting.UseCase
}

// NewEndpoint function to initialize greeting endpoint
func NewEndpoint(cfg config.Config) Endpoint {
	return Endpoint{useCase: usecase.NewUseCase(), cfg: cfg}
}
