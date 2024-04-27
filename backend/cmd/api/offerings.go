package api

import (
	"database/sql"
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/labstack/echo/v4"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"net/http"
)

func (h *Handler) ListOfferings(c echo.Context) error {
	filters := new(schemas.OfferingFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()

	sort, err := filters.SortIn.NewOrderByArrays(schemas.OfferingSortMap)
	if err != nil {
		return err
	}

	filteredOfferings := tbl.Offerings.SELECT(pg.COUNT(tbl.Offerings.ID).OVER().AS("total_records"), tbl.Offerings.AllColumns).
		WHERE(condition).
		ORDER_BY(sort...).
		LIMIT(filters.PaginationIn.Limit()).
		OFFSET(filters.PaginationIn.Offset()).
		AsTable("filtered_offerings")

	sort2, err := filters.SortIn.NewOrderByArrays(schemas.NewOfferingSortMapFrom(filteredOfferings))
	if err != nil {
		return err
	}
	sort2 = append(sort2, tbl.Instructors.Name.ASC())

	stmt := filteredOfferings.INNER_JOIN(tbl.Courses, tbl.Offerings.CourseID.From(filteredOfferings).EQ(tbl.Courses.ID)).
		LEFT_JOIN(tbl.Teaches, tbl.Teaches.OfferingID.EQ(tbl.Offerings.ID.From(filteredOfferings))).
		INNER_JOIN(tbl.Instructors, tbl.Teaches.InstructorID.EQ(tbl.Instructors.ID)).
		SELECT(filteredOfferings.AllColumns(), tbl.Courses.AllColumns, tbl.Instructors.AllColumns).
		ORDER_BY(sort2...)

	var dst []*struct {
		TotalRecords      int64
		FilteredOfferings *model.Offerings
		Courses           *model.Courses
		Instructors       []*model.Instructors
	}

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	var totalRecords int64 = 0
	offerings := make([]*schemas.OfferingOut, len(dst))
	for i, row := range dst {
		totalRecords = row.TotalRecords
		instructors := schemas.NewInstructorOurArray(row.Instructors)
		offerings[i] = schemas.NewOfferingOut(row.FilteredOfferings, schemas.NewCourseOut(row.Courses), instructors)
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(offerings, totalRecords, &filters.PaginationIn))
}

func (h *Handler) CreateOffering(c echo.Context) error {
	input := new(schemas.OfferingIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	tx, err := h.DB.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		if err := tx.Rollback(); err != nil {
			c.Logger().Error(err)
		}
	}(tx)

	stmt1 := tbl.Offerings.INSERT(tbl.Offerings.MutableColumns).
		MODEL(schemas.NewOfferingModel(input)).
		RETURNING(tbl.Offerings.AllColumns)

	var offering model.Offerings
	if err := stmt1.Query(tx, &offering); err != nil {
		return err
	}

	for _, instructorID := range input.InstructorIDs {
		stmt2 := tbl.Teaches.INSERT(tbl.Teaches.AllColumns).
			MODEL(schemas.NewTeachesModel(offering.ID, instructorID))
		if _, err := stmt2.Exec(tx); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, schemas.NewOfferingOutBrief(input, &offering))
}

func (h *Handler) RetrieveOffering(c echo.Context) error {
	input := new(schemas.OfferingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	filteredOffering := tbl.Offerings.SELECT(tbl.Offerings.AllColumns).
		WHERE(tbl.Offerings.ID.EQ(pg.Int64(input.ID))).
		AsTable("filtered_offering")

	filteredOfferingID := tbl.Offerings.ID.From(filteredOffering)
	filteredCourseID := tbl.Offerings.CourseID.From(filteredOffering)

	stmt := filteredOffering.SELECT(filteredOffering.AllColumns(), tbl.Courses.AllColumns, tbl.Instructors.AllColumns).
		FROM(filteredOffering.INNER_JOIN(tbl.Courses, filteredCourseID.EQ(tbl.Courses.ID)).
			LEFT_JOIN(tbl.Teaches, tbl.Teaches.OfferingID.EQ(filteredOfferingID)).
			INNER_JOIN(tbl.Instructors, tbl.Teaches.InstructorID.EQ(tbl.Instructors.ID))).
		ORDER_BY(tbl.Instructors.ID.ASC())

	var dst struct {
		FilteredOffering *model.Offerings
		Courses          *model.Courses
		Instructors      []*model.Instructors
	}

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	instructors := schemas.NewInstructorOurArray(dst.Instructors)

	return c.JSON(http.StatusOK, schemas.NewOfferingOut(dst.FilteredOffering, schemas.NewCourseOut(dst.Courses), instructors))
}

func (h *Handler) UpdateOffering(c echo.Context) error {
	input := new(schemas.OfferingUpdateIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	tx, err := h.DB.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		if err := tx.Rollback(); err != nil {
			c.Logger().Error(err)
		}
	}(tx)

	stmt1 := tbl.Teaches.DELETE().
		WHERE(tbl.Teaches.OfferingID.EQ(pg.Int64(input.ID)))

	if _, err := stmt1.Exec(tx); err != nil {
		return err
	}

	var offering model.Offerings
	stmt2 := tbl.Offerings.UPDATE(tbl.Offerings.MutableColumns).
		MODEL(schemas.NewOfferingModel(&input.OfferingIn)).
		WHERE(tbl.Offerings.ID.EQ(pg.Int64(input.ID))).
		RETURNING(tbl.Offerings.AllColumns)

	if err := stmt2.Query(tx, &offering); err != nil {
		return err
	}

	for _, instructorID := range input.InstructorIDs {
		stmt3 := tbl.Teaches.INSERT(tbl.Teaches.AllColumns).
			MODEL(schemas.NewTeachesModel(offering.ID, instructorID))

		if _, err := stmt3.Exec(tx); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewOfferingOutBrief(&input.OfferingIn, &offering))
}

func (h *Handler) DeleteOffering(c echo.Context) error {
	input := new(schemas.OfferingIDIn)
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := input.Validate(); err != nil {
		return err
	}

	stmt := tbl.Offerings.DELETE().
		WHERE(tbl.Offerings.ID.EQ(pg.Int64(input.ID))).
		RETURNING(tbl.Offerings.ID.AS("id"))

	var dst struct {
		ID int64
	}
	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
