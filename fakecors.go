package fakecors

import "net/http"

func Foo() {}
func Bar() {}
func Baz() {}

type Middleware = func(http.Handler) http.Handler

func NewCORS(opts ...OptionCred) (Middleware, error) {
	return nil, nil
}

type Option interface{}

func ExposeAllResponseHeaders() Option {
	return nil
}

func ExposeResponseHeaders(headers ...string) Option {
	return nil
}

func FromAnyOrigin() Option {
	return nil
}

func FromOrigins(origins ...string) Option {
	return nil
}

func MaxAgeInSeconds(delta uint) Option {
	return nil
}

func PreflightSuccessStatus(code uint) Option {
	return nil
}

func WithAnyMethod() Option {
	return nil
}

func WithAnyRequestHeaders() Option {
	return nil
}

func WithMethods(methods ...string) Option {
	return nil
}

func WithRequestHeaders(headers ...string) Option {
	return nil
}
