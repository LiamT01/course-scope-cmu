package api

import (
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/labstack/echo/v4"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"net/http"
)

func (h *Handler) ListInstructors(c echo.Context) error {
	filters := new(schemas.InstructorFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()

	sort, err := filters.SortIn.NewOrderByArrays(schemas.InstructorSortMap)
	if err != nil {
		return err
	}

	stmt := tbl.Instructors.SELECT(pg.COUNT(tbl.Instructors.ID).OVER().AS("total_records"), tbl.Instructors.AllColumns).
		WHERE(condition).
		ORDER_BY(sort...).
		LIMIT(filters.PaginationIn.Limit()).
		OFFSET(filters.PaginationIn.Offset())

	var totalRecords int64 = 0

	var dst []*struct {
		TotalRecords int64
		*model.Instructors
	}

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	instructors := make([]*schemas.InstructorOut, len(dst))
	for i, row := range dst {
		totalRecords = row.TotalRecords
		instructors[i] = schemas.NewInstructorOut(row.Instructors)
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(instructors, totalRecords, &filters.PaginationIn))
}

func (h *Handler) CreateInstructor(c echo.Context) error {
	instructor := new(schemas.InstructorIn)
	if err := c.Bind(instructor); err != nil {
		return err
	}
	if err := instructor.Validate(); err != nil {
		return err
	}

	src := schemas.NewInstructorModel(instructor)

	stmt := tbl.Instructors.INSERT(tbl.Instructors.MutableColumns).
		MODEL(src).
		RETURNING(tbl.Instructors.AllColumns)

	var dst model.Instructors

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, schemas.NewInstructorOut(&dst))
}

func (h *Handler) RetrieveInstructor(c echo.Context) error {
	instructorIDIn := new(schemas.InstructorIDIn)
	if err := c.Bind(instructorIDIn); err != nil {
		return err
	}
	if err := instructorIDIn.Validate(); err != nil {
		return err
	}

	stmt := tbl.Instructors.SELECT(tbl.Instructors.AllColumns).
		WHERE(tbl.Instructors.ID.EQ(pg.Int64(instructorIDIn.ID)))

	var dst model.Instructors

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewInstructorOut(&dst))
}

func (h *Handler) UpdateInstructor(c echo.Context) error {
	instructor := new(schemas.InstructorUpdateIn)
	if err := c.Bind(instructor); err != nil {
		return err
	}
	if err := instructor.Validate(); err != nil {
		return err
	}

	src := schemas.NewInstructorModel(&instructor.InstructorIn)

	stmt := tbl.Instructors.UPDATE(tbl.Instructors.MutableColumns).
		MODEL(src).
		WHERE(tbl.Instructors.ID.EQ(pg.Int64(instructor.ID))).
		RETURNING(tbl.Instructors.AllColumns)

	var dst model.Instructors

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewInstructorOut(&dst))
}

func (h *Handler) DeleteInstructor(c echo.Context) error {
	instructorIDIn := new(schemas.InstructorIDIn)
	if err := c.Bind(instructorIDIn); err != nil {
		return err
	}
	if err := instructorIDIn.Validate(); err != nil {
		return err
	}

	stmt2 := tbl.Instructors.DELETE().
		WHERE(tbl.Instructors.ID.EQ(pg.Int64(instructorIDIn.ID))).
		RETURNING(tbl.Instructors.ID.AS("id"))

	var dst struct {
		ID int64
	}
	if err := stmt2.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
