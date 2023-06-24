package fakecors

import "net/http"

func Foo() {}
func Bar() {}
func Baz() {}

type Middleware = func(http.Handler) http.Handler

func NewCORS(opts ...func(*config) errorCred) (Middleware, error) {
	return nil, nil
}

type config struct{}

func ExposeAllResponseHeaders() func(*config) error {
	return nil
}

func ExposeResponseHeaders(headers ...string) func(*config) error {
	return nil
}

func FromAnyOrigin() func(*config) error {
	return nil
}

func FromOrigins(origins ...string) func(*config) error {
	return nil
}

func MaxAgeInSeconds(delta uint) func(*config) error {
	return nil
}

func PreflightSuccessStatus(code uint) func(*config) error {
	return nil
}

func WithAnyMethod() func(*config) error {
	return nil
}

func WithAnyRequestHeaders() func(*config) error {
	return nil
}

func WithMethods(methods ...string) func(*config) error {
	return nil
}

func WithRequestHeaders(headers ...string) func(*config) error {
	return nil
}
