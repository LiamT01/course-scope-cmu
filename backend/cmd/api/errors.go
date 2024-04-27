package api

import "errors"

var (
	ErrorTooFrequentRequest = errors.New("too frequent request")
	ErrorInvalidCredentials = errors.New("invalid credentials")
	ErrorInvalidAuthToken   = errors.New("invalid authentication token")
	ErrorUnauthorized       = errors.New("unauthorized")
	ErrorNotActivated       = errors.New("not activated")
	ErrorNotPermitted       = errors.New("not permitted")
	ErrorNotFound           = errors.New("not found")
)
