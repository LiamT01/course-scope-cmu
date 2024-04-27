package api

import (
	"database/sql"
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/iancoleman/strcase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	DB          *sql.DB
	EmailServer *emailServer
}

func NewHandler(db *sql.DB, host, port, hostUser, hostPassword, from, frontendLink string) *Handler {
	return &Handler{
		DB:          db,
		EmailServer: newEmailServer(host, port, hostUser, hostPassword, from, frontendLink),
	}
}

type errorMessageT struct {
	Message string            `json:"message"`
	Detail  map[string]string `json:"detail,omitempty"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	var validationError validation.Errors
	switch {
	case errors.As(err, &validationError):
		errorMessage := createValidationError(validationError)
		if err := c.JSON(http.StatusBadRequest, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	case errors.Is(err, qrm.ErrNoRows) || errors.Is(err, ErrorNotFound):
		errorMessage := errorMessageT{
			Message: "Resource not found.",
		}
		if err := c.JSON(http.StatusNotFound, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	case errors.Is(err, ErrorTooFrequentRequest):
		errorMessage := errorMessageT{
			Message: "You are making too many requests. Please try again later.",
		}
		if err := c.JSON(http.StatusTooManyRequests, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	case errors.Is(err, ErrorInvalidCredentials):
		errorMessage := errorMessageT{
			Message: "Invalid combination of Andrew ID and password.",
		}
		if err := c.JSON(http.StatusUnauthorized, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	case errors.Is(err, ErrorInvalidAuthToken):
		c.Response().Header().Set("WWW-Authenticate", "Bearer")
		errorMessage := errorMessageT{
			Message: "Invalid authentication token.",
		}
		if err := c.JSON(http.StatusUnauthorized, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	case errors.Is(err, ErrorUnauthorized):
		errorMessage := errorMessageT{
			Message: "You must be logged in to perform this action.",
		}
		if err := c.JSON(http.StatusUnauthorized, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	case errors.Is(err, ErrorNotPermitted):
		errorMessage := errorMessageT{
			Message: "You do not have permission to perform this action.",
		}
		if err := c.JSON(http.StatusForbidden, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	case errors.Is(err, ErrorNotActivated):
		errorMessage := errorMessageT{
			Message: "Your account has not been activated.",
		}
		if err := c.JSON(http.StatusForbidden, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	default:
		code := http.StatusInternalServerError
		message := err.Error()
		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code
			message = he.Message.(string)
		}

		errorMessage := errorMessageT{
			Message: message,
		}
		if err := c.JSON(code, errorMessage); err != nil {
			c.Logger().Error(err)
		}
	}
}

func createValidationError(ve validation.Errors) *errorMessageT {
	message := errorMessageT{
		Message: "Validation error.",
		Detail:  make(map[string]string, len(ve)),
	}

	for k, v := range ve {
		message.Detail[strcase.ToSnake(k)] = v.Error()
	}

	return &message
}
