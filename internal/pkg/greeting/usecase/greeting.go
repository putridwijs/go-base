package usecase

import (
	"context"
	"fmt"
	"go-base/internal/pkg/greeting/dto"
)

// Greet function to implement use case
func (u useCase) Greet(ctx context.Context, request dto.GreetingRequest) (dto.GreetingResponse, error) {
	var greeting = "welcome!"
	if request.Name != "" {
		greeting = fmt.Sprintf("hi %s, %s", request.Name, greeting)
	}
	return dto.GreetingResponse{Message: greeting}, nil
}
