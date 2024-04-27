package api

import (
	"errors"
	"fmt"
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/labstack/echo/v4"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"net/http"
)

func (h *Handler) ScanRatings(stmt pg.SelectStatement) ([]*schemas.RatingOut, int64, error) {
	var dst []*struct {
		TotalRecords int64
		model.Ratings
		model.Users
		Offerings *struct {
			model.Offerings
			model.Courses
			Instructors []*model.Instructors
		}
		NetLikes         int64
		LikedByViewer    bool
		DislikedByViewer bool
	}
	if err := stmt.Query(h.DB, &dst); err != nil {
		return nil, 0, err
	}

	var totalRecords int64 = 0
	ratings := make([]*schemas.RatingOut, len(dst))
	for i, row := range dst {
		totalRecords = row.TotalRecords
		user := schemas.NewUserOut(&row.Users)
		course := schemas.NewCourseOut(&row.Offerings.Courses)
		instructors := schemas.NewInstructorOurArray(row.Offerings.Instructors)
		offering := schemas.NewOfferingOut(&row.Offerings.Offerings, course, instructors)
		ratings[i] = schemas.NewRatingOut(&row.Ratings, user, offering, row.NetLikes, row.LikedByViewer, row.DislikedByViewer)
	}

	return ratings, totalRecords, nil
}

func (h *Handler) ListRatings(c echo.Context) error {
	viewer := h.contextGetUser(c)

	filters := new(schemas.RatingFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()

	stmt, err := schemas.FilterAndAnnotateRatingsSQL(&filters.SortIn, &filters.PaginationIn, condition, viewer.ID)
	if err != nil {
		return err
	}

	ratings, totalRecords, err := h.ScanRatings(stmt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(ratings, totalRecords, &filters.PaginationIn))
}

// ListMyRatings requires auth token
func (h *Handler) ListMyRatings(c echo.Context) error {
	viewer := h.contextGetUser(c)

	filters := new(schemas.RatingFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()
	condition = condition.AND(tbl.Ratings.UserID.EQ(pg.Int64(viewer.ID)))

	stmt, err := schemas.FilterAndAnnotateRatingsSQL(&filters.SortIn, &filters.PaginationIn, condition, viewer.ID)
	if err != nil {
		return err
	}

	ratings, totalRecords, err := h.ScanRatings(stmt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(ratings, totalRecords, &filters.PaginationIn))
}

func (h *Handler) ListMyLikedRatings(c echo.Context) error {
	viewer := h.contextGetUser(c)

	filters := new(schemas.RatingFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()
	condition = condition.AND(pg.EXISTS(
		tbl.Likes.SELECT(tbl.Likes.ID).
			WHERE(tbl.Likes.RatingID.EQ(tbl.Ratings.ID).
				AND(tbl.Likes.UserID.EQ(pg.Int64(viewer.ID)))),
	))

	stmt, err := schemas.FilterAndAnnotateRatingsSQL(&filters.SortIn, &filters.PaginationIn, condition, viewer.ID)
	if err != nil {
		return err
	}

	ratings, totalRecords, err := h.ScanRatings(stmt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(ratings, totalRecords, &filters.PaginationIn))
}

func (h *Handler) ListMyDislikedRatings(c echo.Context) error {
	viewer := h.contextGetUser(c)

	filters := new(schemas.RatingFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()
	condition = condition.AND(pg.EXISTS(
		tbl.Dislikes.SELECT(tbl.Dislikes.ID).
			WHERE(tbl.Dislikes.RatingID.EQ(tbl.Ratings.ID).
				AND(tbl.Dislikes.UserID.EQ(pg.Int64(viewer.ID)))),
	))

	stmt, err := schemas.FilterAndAnnotateRatingsSQL(&filters.SortIn, &filters.PaginationIn, condition, viewer.ID)
	if err != nil {
		return err
	}

	ratings, totalRecords, err := h.ScanRatings(stmt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(ratings, totalRecords, &filters.PaginationIn))
}

func (h *Handler) CreateRating(c echo.Context) error {
	input := new(schemas.RatingIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	user := h.contextGetUser(c)
	input.UserID = user.ID

	stmt := tbl.Ratings.
		INSERT(
			tbl.Ratings.UserID,
			tbl.Ratings.OfferingID,
			tbl.Ratings.Overall,
			tbl.Ratings.Teaching,
			tbl.Ratings.Materials,
			tbl.Ratings.Value,
			tbl.Ratings.Difficulty,
			tbl.Ratings.Workload,
			tbl.Ratings.Grading,
			tbl.Ratings.Comment,
		).
		MODEL(schemas.NewRatingModel(input)).
		RETURNING(tbl.Ratings.AllColumns)

	var dst model.Ratings
	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, schemas.NewRatingOutBrief(input, &dst))
}

func (h *Handler) RetrieveRating(c echo.Context) error {
	viewer := h.contextGetUser(c)

	input := new(schemas.RatingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt, err := schemas.FilterAndAnnotateRatingsSQL(nil, nil, tbl.Ratings.ID.EQ(pg.Int64(input.ID)), viewer.ID)
	if err != nil {
		return err
	}

	ratings, totalRecords, err := h.ScanRatings(stmt)
	if err != nil {
		return err
	}
	if totalRecords == 0 {
		return ErrorNotFound
	}

	dst := ratings[0]

	return c.JSON(http.StatusOK, dst)
}

func (h *Handler) UpdateRating(c echo.Context) error {
	input := new(schemas.RatingUpdateIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	user := h.contextGetUser(c)

	stmt := tbl.Ratings.
		UPDATE(
			tbl.Ratings.OfferingID,
			tbl.Ratings.UpdatedAt,
			tbl.Ratings.Overall,
			tbl.Ratings.Teaching,
			tbl.Ratings.Materials,
			tbl.Ratings.Value,
			tbl.Ratings.Difficulty,
			tbl.Ratings.Workload,
			tbl.Ratings.Grading,
			tbl.Ratings.Comment,
		).
		MODEL(schemas.NewRatingModel(&input.RatingIn)).
		WHERE(tbl.Ratings.ID.EQ(pg.Int64(input.ID)).
			AND(tbl.Ratings.UserID.EQ(pg.Int64(user.ID)))).
		RETURNING(tbl.Ratings.AllColumns)

	var dst model.Ratings
	if err := stmt.Query(h.DB, &dst); err != nil {
		if dst.UserID != user.ID {
			return ErrorNotPermitted
		}
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewRatingOutBrief(&input.RatingIn, &dst))
}

func (h *Handler) DeleteRating(c echo.Context) error {
	input := new(schemas.RatingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	user := h.contextGetUser(c)

	stmt := tbl.Ratings.DELETE().
		WHERE(tbl.Ratings.ID.EQ(pg.Int64(input.ID)).
			AND(tbl.Ratings.UserID.EQ(pg.Int64(user.ID)))).
		RETURNING(tbl.Ratings.ID.AS("id"), tbl.Ratings.UserID.AS("user_id"))

	var dst struct {
		ID     int64
		UserID int64
	}
	if err := stmt.Query(h.DB, &dst); err != nil {
		if dst.UserID != user.ID {
			return ErrorNotPermitted
		}
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) LikeRating(c echo.Context) error {
	viewer := h.contextGetUser(c)

	input := new(schemas.RatingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt := pg.SELECT(pg.EXISTS(
		tbl.Likes.SELECT(tbl.Likes.ID).
			WHERE(tbl.Likes.RatingID.EQ(pg.Int64(input.ID)).
				AND(tbl.Likes.UserID.EQ(pg.Int64(viewer.ID)))),
	).AS("exists"))
	var exists struct {
		Exists bool
	}
	if err := stmt.Query(h.DB, &exists); err != nil {
		return err
	}
	if exists.Exists {
		return echo.NewHTTPError(http.StatusConflict, "You have already liked this rating.")
	}

	stmt2 := tbl.Likes.INSERT(tbl.Likes.RatingID, tbl.Likes.UserID).
		MODEL(schemas.NewLikeModel(input.ID, viewer.ID)).
		RETURNING(tbl.Likes.AllColumns)
	var dst model.Likes

	if err := stmt2.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, schemas.NewLikeOut(&dst))
}

func (h *Handler) UndoLikeRating(c echo.Context) error {
	viewer := h.contextGetUser(c)

	input := new(schemas.RatingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt := tbl.Likes.DELETE().
		WHERE(tbl.Likes.RatingID.EQ(pg.Int64(input.ID)).
			AND(tbl.Likes.UserID.EQ(pg.Int64(viewer.ID)))).
		RETURNING(tbl.Likes.ID.AS("id"))

	var dst struct {
		ID int64
	}
	if err := stmt.Query(h.DB, &dst); err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return echo.NewHTTPError(http.StatusConflict, "You have not liked this rating.")
		}
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DislikeRating(c echo.Context) error {
	viewer := h.contextGetUser(c)

	input := new(schemas.RatingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt := pg.SELECT(pg.EXISTS(
		tbl.Dislikes.SELECT(tbl.Dislikes.ID).
			WHERE(tbl.Dislikes.RatingID.EQ(pg.Int64(input.ID)).
				AND(tbl.Dislikes.UserID.EQ(pg.Int64(viewer.ID)))),
	).AS("exists"))
	var exists struct {
		Exists bool
	}
	if err := stmt.Query(h.DB, &exists); err != nil {
		return err
	}
	if exists.Exists {
		return echo.NewHTTPError(http.StatusConflict, "You have already disliked this rating.")
	}

	stmt2 := tbl.Dislikes.INSERT(tbl.Dislikes.RatingID, tbl.Dislikes.UserID).
		MODEL(schemas.NewDislikeModel(input.ID, viewer.ID)).
		RETURNING(tbl.Dislikes.AllColumns)
	var dst model.Dislikes

	if err := stmt2.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, schemas.NewDislikeOut(&dst))
}

func (h *Handler) UndoDislikeRating(c echo.Context) error {
	viewer := h.contextGetUser(c)

	input := new(schemas.RatingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt := tbl.Dislikes.DELETE().
		WHERE(tbl.Dislikes.RatingID.EQ(pg.Int64(input.ID)).
			AND(tbl.Dislikes.UserID.EQ(pg.Int64(viewer.ID)))).
		RETURNING(tbl.Dislikes.ID.AS("id"))

	var dst struct {
		ID int64
	}
	if err := stmt.Query(h.DB, &dst); err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return echo.NewHTTPError(http.StatusConflict, "You have not disliked this rating.")
		}
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) CalculateRatingStats(c echo.Context) error {
	filters := new(schemas.RatingFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()
	filteredRatings, err := schemas.FilterRatingsSQL(nil, nil, condition, 0)
	if err != nil {
		return err
	}
	filteredRatingsTable := filteredRatings.AsTable("filtered_ratings")

	columnMap := map[string]pg.ColumnInteger{
		"overall":    tbl.Ratings.Overall.From(filteredRatingsTable),
		"teaching":   tbl.Ratings.Teaching.From(filteredRatingsTable),
		"materials":  tbl.Ratings.Materials.From(filteredRatingsTable),
		"value":      tbl.Ratings.Value.From(filteredRatingsTable),
		"difficulty": tbl.Ratings.Difficulty.From(filteredRatingsTable),
		"workload":   tbl.Ratings.Workload.From(filteredRatingsTable),
		"grading":    tbl.Ratings.Grading.From(filteredRatingsTable),
	}

	var projections []pg.Projection

	// Append average columns
	for k, v := range columnMap {
		projections = append(projections, pg.AVG(v).AS(fmt.Sprintf("rating_stats_model.avg_%s", k)))
	}

	// Append count columns for each value of each rating rubric column
	for k, v := range columnMap {
		for i := int32(1); i <= int32(5); i++ {
			projections = append(projections, pg.COUNT(pg.CASE().WHEN(v.EQ(pg.Int32(i))).THEN(pg.Int32(1))).
				AS(fmt.Sprintf("rating_stats_model.%s_%d", k, i)))
		}
	}

	stmt := filteredRatingsTable.SELECT(
		pg.COUNT(tbl.Ratings.ID.From(filteredRatingsTable)).AS("rating_stats_model.rating_count"),
		projections...,
	)

	var dst schemas.RatingStatsModel
	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewRatingStatsOut(&dst))
}
