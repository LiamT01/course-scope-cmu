package api

import (
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/labstack/echo/v4"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"net/http"
	"time"
)

func (h *Handler) ListUsers(c echo.Context) error {
	filters := new(schemas.UserFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	sort, err := filters.SortIn.NewOrderByArrays(schemas.UserSortMap)
	if err != nil {
		return err
	}

	stmt := tbl.Users.SELECT(
		pg.COUNT(tbl.Users.ID).OVER().AS("total_records"),
		tbl.Users.ID,
		tbl.Users.Username,
		tbl.Users.Activated,
		tbl.Users.CreatedAt,
	).
		ORDER_BY(sort...).
		LIMIT(filters.PaginationIn.Limit()).
		OFFSET(filters.PaginationIn.Offset())

	var totalRecord int64 = 0

	var dst []struct {
		TotalRecords int64
		model.Users
	}

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	users := make([]*schemas.UserOut, len(dst))
	for i, row := range dst {
		totalRecord = row.TotalRecords
		users[i] = schemas.NewUserOut(&row.Users)
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(users, totalRecord, &filters.PaginationIn))
}

func (h *Handler) CreateUser(c echo.Context) error {
	userIn := new(schemas.UserIn)
	if err := c.Bind(userIn); err != nil {
		return err
	}
	if err := userIn.Validate(); err != nil {
		return err
	}

	var exists struct {
		Exists bool
	}

	stmt0 := tbl.Users.SELECT(tbl.Users.ID).WHERE(pg.LOWER(tbl.Users.AndrewID).EQ(pg.LOWER(pg.String(userIn.AndrewID))))
	stmt1 := pg.SELECT(pg.EXISTS(stmt0).AS("exists"))
	if err := stmt1.Query(h.DB, &exists); err != nil {
		return err
	}
	if exists.Exists {
		return echo.NewHTTPError(http.StatusConflict, "This Andrew ID is already registered. Please log in or reset your password.")
	}

	stmt0 = tbl.Users.SELECT(tbl.Users.ID).WHERE(tbl.Users.Username.EQ(pg.String(userIn.Username)))
	stmt1 = pg.SELECT(pg.EXISTS(stmt0).AS("exists"))
	if err := stmt1.Query(h.DB, &exists); err != nil {
		return err
	}
	if exists.Exists {
		return echo.NewHTTPError(http.StatusConflict, "This username is already taken. Please choose another one.")
	}

	src, err := schemas.NewUserModel(userIn)
	if err != nil {
		return err
	}

	stmt2 := tbl.Users.INSERT(tbl.Users.AndrewID, tbl.Users.Username, tbl.Users.PasswordHash).
		MODEL(src).
		RETURNING(tbl.Users.AllColumns)

	var dst struct {
		model.Users
	}
	if err := stmt2.Query(h.DB, &dst); err != nil {
		return err
	}

	tokenOut, err := schemas.NewTokenOut(dst.Users.ID, 1*time.Hour, model.ScopeType_Act)
	if err != nil {
		return err
	}

	if err := h.InsertToken(tokenOut); err != nil {
		return err
	}

	if err := h.EmailServer.sendTokenToUser(&dst.Users, tokenOut); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, schemas.NewUserOut(&dst.Users))
}

func (h *Handler) RetrieveUser(c echo.Context) error {
	input := new(schemas.UserIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt := tbl.Users.SELECT(tbl.Users.ID, tbl.Users.Username, tbl.Users.Activated, tbl.Users.CreatedAt).
		WHERE(tbl.Users.ID.EQ(pg.Int64(input.ID)))

	var dst struct {
		model.Users
	}

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewUserOut(&dst.Users))
}

// RetrieveUserMe requires auth token
func (h *Handler) RetrieveUserMe(c echo.Context) error {
	user := h.contextGetUser(c)
	return c.JSON(http.StatusOK, schemas.NewUserOut(user))
}

func (h *Handler) UpdateUser(c echo.Context) error {
	input := new(schemas.UserUpdateIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	src, err := schemas.NewUserModel(&input.UserIn)
	if err != nil {
		return err
	}

	stmt := tbl.Users.UPDATE(tbl.Users.AndrewID, tbl.Users.Username, tbl.Users.PasswordHash).
		MODEL(src).
		WHERE(tbl.Users.ID.EQ(pg.Int64(input.ID))).
		RETURNING(tbl.Users.AllColumns)

	var dst struct {
		model.Users
	}
	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewUserOut(&dst.Users))
}

func (h *Handler) DeleteUser(c echo.Context) error {
	user := h.contextGetUser(c)

	input := new(schemas.UserIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	if user.ID != input.ID {
		return ErrorNotPermitted
	}

	stmt := tbl.Users.DELETE().
		WHERE(tbl.Users.ID.EQ(pg.Int64(input.ID))).
		RETURNING(tbl.Users.ID.AS("id"))

	var dst struct {
		ID int64
	}
	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ActivateUser(c echo.Context) error {
	input := new(schemas.TokenIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	user, err := h.RetrieveUserForValidToken(model.ScopeType_Act, input.Token)
	if err != nil {
		return err
	}

	user.Activated = true
	stmt := tbl.Users.UPDATE(tbl.Users.Activated).
		MODEL(user).
		WHERE(tbl.Users.ID.EQ(pg.Int64(user.ID)))
	if _, err := stmt.Exec(h.DB); err != nil {
		return err
	}

	if err := h.DeleteAllTokensForUser(user.ID, model.ScopeType_Act); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewUserOut(user))
}

func (h *Handler) UpdateUserPassword(c echo.Context) error {
	input := new(schemas.ResetPasswordIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	user, err := h.RetrieveUserForValidToken(model.ScopeType_Pwd, input.Token)
	if err != nil {
		return err
	}

	user.PasswordHash, err = schemas.HashPasswordPlaintext(input.Plaintext)
	if err != nil {
		return err
	}

	stmt := tbl.Users.UPDATE(tbl.Users.PasswordHash).
		MODEL(user).
		WHERE(tbl.Users.ID.EQ(pg.Int64(user.ID)))
	if _, err := stmt.Exec(h.DB); err != nil {
		return err
	}

	if err := h.DeleteAllTokensForUser(user.ID, model.ScopeType_Pwd); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewUserOut(user))
}

// UpdateUsername requires auth token that is activated
func (h *Handler) UpdateUsername(c echo.Context) error {
	input := new(schemas.UsernameIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt1 := pg.SELECT(pg.EXISTS(tbl.Users.SELECT(tbl.Users.ID).
		WHERE(tbl.Users.Username.EQ(pg.String(input.Username)))).AS("exists"))
	var dst struct {
		Exists bool
	}
	if err := stmt1.Query(h.DB, &dst); err != nil {
		return err
	}
	if dst.Exists {
		return echo.NewHTTPError(http.StatusConflict, "This username is already taken. Please choose another one.")
	}

	user := h.contextGetUser(c)
	user.Username = input.Username
	stmt2 := tbl.Users.UPDATE(tbl.Users.Username).
		MODEL(user).
		WHERE(tbl.Users.ID.EQ(pg.Int64(user.ID)))
	if _, err := stmt2.Exec(h.DB); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewUserOut(user))
}

func (h *Handler) CalculateUserStats(c echo.Context) error {
	user := h.contextGetUser(c)

	stmt := pg.SELECT(
		tbl.Likes.INNER_JOIN(tbl.Ratings, tbl.Ratings.ID.EQ(tbl.Likes.RatingID)).
			SELECT(pg.COUNT(tbl.Likes.ID)).
			WHERE(tbl.Ratings.UserID.EQ(pg.Int64(user.ID))).
			AS("likes_received"),
		tbl.Ratings.INNER_JOIN(tbl.Offerings, tbl.Offerings.ID.EQ(tbl.Ratings.OfferingID)).
			INNER_JOIN(tbl.Courses, tbl.Courses.ID.EQ(tbl.Offerings.CourseID)).
			SELECT(pg.COUNT(pg.DISTINCT(tbl.Courses.ID))).
			WHERE(tbl.Ratings.UserID.EQ(pg.Int64(user.ID))).
			AS("courses_rated"),
	)
	var dst struct {
		LikesReceived int64
		CoursesRated  int64
	}

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewUserStatsOut(dst.LikesReceived, dst.CoursesRated))
}
