package utils

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Echo utility instance
var Echo EchoUtil

// EchoBinder utility instance
var EchoBinder EchoBinderWithValidation

// EchoUtil is a utility struct for working with Echo instances
type EchoUtil struct{}

// EchoValidator is a struct that implements the echo.Validator interface.
type EchoValidator struct{}

// Validate is a method that validates the given struct using the ValidateStruct
// function and returns an error if validation fails.
func (EchoValidator) Validate(i any) error {
	return ValidateStruct(i)
}

// EchoBinderWithValidation is a struct that implements the echo.Binder interface
// with added validation functionality.
type EchoBinderWithValidation struct {
	echo.DefaultBinder
}

// Bind is a method that binds the request data to the given struct,
// validates it using the ValidateStruct function,
// and returns an error if binding or validation fails.
func (b *EchoBinderWithValidation) Bind(i any, c echo.Context) error {
	if err := b.DefaultBinder.Bind(i, c); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}
	return ValidateStruct(i)
}

// BindBody is a method that binds the body data to the given struct,
// validates it using the ValidateStruct function,
// and returns an error if binding or validation fails.
func (b *EchoBinderWithValidation) BindBody(c echo.Context, i any) (err error) {
	if err := b.DefaultBinder.BindBody(c, i); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}
	return ValidateStruct(i)
}

// BindHeaders is a method that binds the headers data to the given struct,
// validates it using the ValidateStruct function,
// and returns an error if binding or validation fails.
func (b *EchoBinderWithValidation) BindHeaders(c echo.Context, i any) (err error) {
	if err := b.DefaultBinder.BindHeaders(c, i); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}
	return ValidateStruct(i)
}

// BindPathParams is a method that binds the path params to the given struct,
// validates it using the ValidateStruct function,
// and returns an error if binding or validation fails.
func (b *EchoBinderWithValidation) BindPathParams(c echo.Context, i any) (err error) {
	if err := b.DefaultBinder.BindPathParams(c, i); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}
	return ValidateStruct(i)
}

// BindQueryParams is a method that binds the query params to the given struct,
// validates it using the ValidateStruct function,
// and returns an error if binding or validation fails.
func (b *EchoBinderWithValidation) BindQueryParams(c echo.Context, i any) (err error) {
	if err := b.DefaultBinder.BindQueryParams(c, i); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}
	return ValidateStruct(i)
}

// DefaultRootHandler handles requests to the root endpoint
func (EchoUtil) DefaultRootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "200 OK"})
}

// NoContentHandler handles return no content endpoint
func (EchoUtil) NoContentHandler(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// New creates a new instance of the Echo framework
func (EchoUtil) New() *echo.Echo {
	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Validator = new(EchoValidator)
	e.Binder = &EchoBinder
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/", Echo.DefaultRootHandler)
	e.GET("/favicon.ico", Echo.NoContentHandler)
	return e
}
