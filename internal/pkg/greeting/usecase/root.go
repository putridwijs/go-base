package usecase

import "go-base/internal/pkg/greeting"

type useCase struct {
}

// NewUseCase function to initialize greeting use case
func NewUseCase() greeting.UseCase {
	return useCase{}
}
