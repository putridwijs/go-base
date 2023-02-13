package greeting

import "github.com/go-kit/kit/endpoint"

// Endpoint interface of greeting package
type Endpoint interface {
	Greet() endpoint.Endpoint
}
