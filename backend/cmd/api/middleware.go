package api

import (
	"context"
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/labstack/echo/v4"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"strings"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const userContextKey = contextKey("user")

func (h *Handler) contextSetUser(c echo.Context, user *model.Users) {
	c.SetRequest(c.Request().WithContext(
		context.WithValue(c.Request().Context(), userContextKey, user)))
}

func (h *Handler) contextGetUser(c echo.Context) *model.Users {
	user, ok := c.Request().Context().Value(userContextKey).(*model.Users)
	if !ok {
		panic("missing user in request context")
	}
	return user
}

func (h *Handler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			h.contextSetUser(c, schemas.AnonymousUser)
			return next(c)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return ErrorInvalidAuthToken
		}

		token := headerParts[1]
		tokenIn := &schemas.TokenIn{Token: token}
		if err := tokenIn.Validate(); err != nil {
			return ErrorInvalidAuthToken
		}

		user, err := h.RetrieveUserForValidToken(model.ScopeType_Auth, token)
		if err != nil {
			switch {
			case errors.Is(err, qrm.ErrNoRows):
				return ErrorInvalidAuthToken
			default:
				return err
			}
		}

		h.contextSetUser(c, user)

		return next(c)
	}
}

func (h *Handler) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := h.contextGetUser(c)
		if schemas.IsAnonymousUser(user) {
			return ErrorUnauthorized
		}
		return next(c)
	}
}

func (h *Handler) RequireActivatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := h.contextGetUser(c)
		if schemas.IsAnonymousUser(user) {
			return ErrorUnauthorized
		}
		if !schemas.IsActivatedUser(user) {
			return ErrorNotActivated
		}
		return next(c)
	}
}

func (h *Handler) RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := h.contextGetUser(c)
		if !schemas.IsAdminUser(user) {
			return ErrorNotPermitted
		}
		return next(c)
	}
}
