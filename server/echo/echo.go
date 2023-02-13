package echo

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/labstack/echo/v4"
)

// Handler is function to adapt go-kit to echo http framework
func Handler(endpoint endpoint.Endpoint, opts ...Option) func(ctx echo.Context) error {
	var settings = defaultOptions
	for _, opt := range opts {
		opt.Apply(&settings)
	}

	return func(ctx echo.Context) error {
		errFunc := func(err error) error {
			if settings.errorHandler != nil {
				return settings.errorHandler(ctx, err)
			}
			return err
		}

		request, err := settings.decoder(ctx)
		if err != nil {
			return errFunc(err)
		}

		response, err := endpoint(ctx.Request().Context(), request)
		if err != nil {
			return errFunc(err)
		}

		return settings.encoder(ctx, response)
	}
}
