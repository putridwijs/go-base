package encoding

import (
	"github.com/labstack/echo/v4"
	"go-base/internal/pkg/greeting/dto"
)

func DecodeGreetingRequest() func(ctx echo.Context) (interface{}, error) {
	return func(ctx echo.Context) (interface{}, error) {
		var request dto.GreetingRequest
		if err := ctx.Bind(&request); err != nil {
			return nil, err
		}

		return request, nil
	}
}
