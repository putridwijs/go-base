package echo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Option interface {
		Apply(o *options)
	}
	options struct {
		encoder      func(ctx echo.Context, response interface{}) error
		decoder      func(ctx echo.Context) (interface{}, error)
		errorHandler func(ctx echo.Context, err error) error
	}
)

var (
	defaultDecoder = func(ctx echo.Context) (interface{}, error) {
		return nil, nil
	}
	defaultEncoder = func(ctx echo.Context, response interface{}) error {
		return ctx.JSON(http.StatusOK, response)
	}
	defaultOptions = options{
		encoder:      defaultEncoder,
		decoder:      defaultDecoder,
		errorHandler: nil,
	}
)

// Decoder Option
type withDecoder func(ctx echo.Context) (interface{}, error)

func (w withDecoder) Apply(o *options) {
	o.decoder = w
}

func WithDecoder(decoder func(ctx echo.Context) (interface{}, error)) Option {
	return withDecoder(decoder)
}

// Encoder Option
type withEncoder func(ctx echo.Context, response interface{}) error

func (w withEncoder) Apply(o *options) {
	o.encoder = w
}

func WithEncoder(encoder func(ctx echo.Context, response interface{}) error) Option {
	return withEncoder(encoder)
}

// ErrorHandler Option
type withErrorHandler func(ctx echo.Context, err error) error

func (w withErrorHandler) Apply(o *options) {
	o.errorHandler = w
}

func WithErrorHandler(errorHandler func(ctx echo.Context, err error) error) Option {
	return withErrorHandler(errorHandler)
}
