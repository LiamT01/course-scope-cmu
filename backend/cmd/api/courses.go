package api

import (
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/labstack/echo/v4"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"net/http"
)

func (h *Handler) ListCourses(c echo.Context) error {
	filters := new(schemas.CourseFilters)
	if err := c.Bind(filters); err != nil {
		return err
	}
	if err := filters.Validate(); err != nil {
		return err
	}

	condition := filters.NewFilterCondition()

	sort, err := filters.SortIn.NewOrderByArrays(schemas.CourseSortMap)
	if err != nil {
		return err
	}

	stmt :=
		tbl.Courses.SELECT(pg.COUNT(tbl.Courses.ID).OVER().AS("total_records"), tbl.Courses.AllColumns).
			WHERE(condition).
			ORDER_BY(sort...).
			LIMIT(filters.PaginationIn.Limit()).
			OFFSET(filters.PaginationIn.Offset())

	var totalRecords int64 = 0

	var dst []*struct {
		TotalRecords int64
		*model.Courses
	}

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	courses := make([]*schemas.CourseOut, len(dst))
	for i, row := range dst {
		totalRecords = row.TotalRecords
		courses[i] = schemas.NewCourseOut(row.Courses)
	}

	return c.JSON(http.StatusOK, schemas.NewEnvelope(courses, totalRecords, &filters.PaginationIn))
}

func (h *Handler) CreateCourse(c echo.Context) error {
	course := new(schemas.CourseIn)
	if err := c.Bind(course); err != nil {
		return err
	}
	if err := course.Validate(); err != nil {
		return err
	}

	src := schemas.NewCourseModel(course)

	stmt := tbl.Courses.INSERT(tbl.Courses.MutableColumns).
		MODEL(src).
		RETURNING(tbl.Courses.AllColumns)

	var dst model.Courses

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, schemas.NewCourseOut(&dst))
}

func (h *Handler) RetrieveCourse(c echo.Context) error {
	courseIDIn := new(schemas.CourseIDIn)
	if err := c.Bind(courseIDIn); err != nil {
		return err
	}
	if err := courseIDIn.Validate(); err != nil {
		return err
	}

	stmt := tbl.Courses.SELECT(tbl.Courses.AllColumns).
		WHERE(tbl.Courses.ID.EQ(pg.Int64(courseIDIn.ID)))

	var dst model.Courses

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewCourseOut(&dst))
}

func (h *Handler) UpdateCourse(c echo.Context) error {
	courseUpdateIn := new(schemas.CourseUpdateIn)
	if err := c.Bind(courseUpdateIn); err != nil {
		return err
	}
	if err := courseUpdateIn.Validate(); err != nil {
		return err
	}

	src := schemas.NewCourseModel(&courseUpdateIn.CourseIn)

	stmt := tbl.Courses.UPDATE(tbl.Courses.MutableColumns).
		MODEL(src).
		WHERE(tbl.Courses.ID.EQ(pg.Int64(courseUpdateIn.ID))).
		RETURNING(tbl.Courses.AllColumns)

	var dst model.Courses

	if err := stmt.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemas.NewCourseOut(&dst))
}

func (h *Handler) DeleteCourse(c echo.Context) error {
	courseIDIn := new(schemas.CourseIDIn)
	if err := c.Bind(courseIDIn); err != nil {
		return err
	}
	if err := courseIDIn.Validate(); err != nil {
		return err
	}

	stmt2 := tbl.Courses.DELETE().
		WHERE(tbl.Courses.ID.EQ(pg.Int64(courseIDIn.ID))).
		RETURNING(tbl.Courses.ID.AS("id"))

	var dst struct {
		ID int64
	}
	if err := stmt2.Query(h.DB, &dst); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
