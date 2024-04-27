package api

import (
	"errors"
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/labstack/echo/v4"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"net/http"
	"time"
)

// RetrieveUserForValidToken retrieves a user (may be nil) for a valid token.
func (h *Handler) RetrieveUserForValidToken(scope model.ScopeType, plaintext string) (*model.Users, error) {
	hashString := schemas.HashTokenPlaintext(plaintext)

	stmt := tbl.Users.INNER_JOIN(tbl.Tokens, tbl.Users.ID.EQ(tbl.Tokens.UserID)).
		SELECT(tbl.Users.AllColumns).
		WHERE(tbl.Tokens.Hash.EQ(pg.Bytea(hashString)).
			AND(pg.CAST(tbl.Tokens.Scope).AS_TEXT().EQ(pg.String(scope.String()))).
			AND(tbl.Tokens.Expiry.GT(pg.TimestampzT(time.Now()))),
		)

	var dst model.Users
	if err := stmt.Query(h.DB, &dst); err != nil {
		return nil, err
	}

	return &dst, nil
}

func (h *Handler) InsertToken(token *schemas.TokenOut) error {
	stmt1 := pg.SELECT(pg.EXISTS(tbl.Tokens.SELECT(tbl.Tokens.ID).
		WHERE(tbl.Tokens.UserID.EQ(pg.Int64(token.UserID)).
			AND(pg.CAST(tbl.Tokens.Scope).AS_TEXT().EQ(pg.String(token.Scope.String()))).
			AND(tbl.Tokens.Expiry.GT(pg.TimestampzT(time.Now()))).
			AND(tbl.Tokens.CreatedAt.GT(pg.TimestampzT(time.Now().Add(-30 * time.Second)))),
		)).AS("exists"))
	var dst struct {
		Exists bool
	}
	if err := stmt1.Query(h.DB, &dst); err != nil {
		return err
	}
	if dst.Exists {
		return ErrorTooFrequentRequest
	}

	stmt2 := tbl.Tokens.INSERT(tbl.Tokens.Hash, tbl.Tokens.UserID, tbl.Tokens.Expiry, tbl.Tokens.Scope).
		MODEL(schemas.NewTokenModel(token))
	if _, err := stmt2.Exec(h.DB); err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteAllTokensForUser(userID int64, scope model.ScopeType) error {
	stmt := tbl.Tokens.DELETE().
		WHERE(tbl.Tokens.UserID.EQ(pg.Int64(userID)).
			AND(pg.CAST(tbl.Tokens.Scope).AS_TEXT().EQ(pg.String(scope.String()))),
		)
	_, err := stmt.Exec(h.DB)
	return err
}

func (h *Handler) DeleteExpiredTokensForUser(userID int64) error {
	stmt := tbl.Tokens.DELETE().
		WHERE(tbl.Tokens.UserID.EQ(pg.Int64(userID)).
			AND(tbl.Tokens.Expiry.LT_EQ(pg.TimestampzT(time.Now()))),
		)
	_, err := stmt.Exec(h.DB)
	return err
}

// AuthenticateUser returns the user if the credentials are valid, or returns nil with an error.
func (h *Handler) AuthenticateUser(andrewID, pwdPlaintext string) (*model.Users, error) {
	stmt := tbl.Users.SELECT(tbl.Users.AllColumns).
		WHERE(tbl.Users.AndrewID.EQ(pg.String(andrewID)))
	var dst model.Users
	if err := stmt.Query(h.DB, &dst); err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, ErrorInvalidCredentials
		}
		return nil, err
	}

	ok, err := schemas.CheckPassword(pwdPlaintext, dst.PasswordHash)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrorInvalidCredentials
	}

	return &dst, nil
}

// GetContextUserOrAuthUser returns the user from the context or credentials, or returns nil with an error.
func (h *Handler) GetContextUserOrAuthUser(c echo.Context) (*model.Users, error) {
	user := h.contextGetUser(c)

	if !schemas.IsAnonymousUser(user) {
		return user, nil
	}

	input := new(schemas.CredentialsIn)
	if err := c.Bind(input); err != nil {
		return nil, err
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user, err := h.AuthenticateUser(input.AndrewID, input.Plaintext)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateAuthToken requires NO token
func (h *Handler) CreateAuthToken(c echo.Context) error {
	credentials := new(schemas.CredentialsIn)
	if err := c.Bind(credentials); err != nil {
		return err
	}
	if err := credentials.Validate(); err != nil {
		return err
	}

	user, err := h.AuthenticateUser(credentials.AndrewID, credentials.Plaintext)
	if err != nil {
		return err
	}

	if err := h.DeleteExpiredTokensForUser(user.ID); err != nil {
		return err
	}

	tokenOut, err := schemas.NewTokenOut(user.ID, 15*24*time.Hour, model.ScopeType_Auth)
	if err != nil {
		return err
	}

	if err := h.InsertToken(tokenOut); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, tokenOut)
}

// DeleteAllAuthTokensForUser requires no token
func (h *Handler) DeleteAllAuthTokensForUser(c echo.Context) error {
	user, err := h.GetContextUserOrAuthUser(c)
	if err != nil {
		return err
	}

	if err := h.DeleteAllTokensForUser(user.ID, model.ScopeType_Auth); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// SendActToken requires auth token
func (h *Handler) SendActToken(c echo.Context) error {
	user := h.contextGetUser(c)
	if schemas.IsActivatedUser(user) {
		return echo.NewHTTPError(http.StatusConflict, "Your account is already activated.")
	}

	tokenOut, err := schemas.NewTokenOut(user.ID, 1*time.Hour, model.ScopeType_Act)
	if err != nil {
		return err
	}
	if err := h.InsertToken(tokenOut); err != nil {
		return err
	}

	if err := h.EmailServer.sendTokenToUser(user, tokenOut); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Please check your email for the activation link.",
	})
}

// SendPwdResetToken requires NO token
func (h *Handler) SendPwdResetToken(c echo.Context) error {
	user := h.contextGetUser(c)
	if schemas.IsAnonymousUser(user) {
		input := new(schemas.AndrewIDIn)
		if err := c.Bind(input); err != nil {
			return err
		}
		if err := input.Validate(); err != nil {
			return err
		}

		stmt := tbl.Users.SELECT(tbl.Users.AllColumns).
			WHERE(tbl.Users.AndrewID.EQ(pg.String(input.AndrewID)))
		var dst model.Users
		if err := stmt.Query(h.DB, &dst); err != nil {
			if errors.Is(err, qrm.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, "The provided Andrew ID does not exist.")
			}
			return err
		}

		user = &dst
	}

	tokenOut, err := schemas.NewTokenOut(user.ID, 1*time.Hour, model.ScopeType_Pwd)
	if err != nil {
		return err
	}

	if err := h.InsertToken(tokenOut); err != nil {
		return err
	}

	if err := h.EmailServer.sendTokenToUser(user, tokenOut); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Please check your email for the password reset link.",
	})
}
