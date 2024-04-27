package schemas

import (
	"github.com/dlclark/regexp2"
	pg "github.com/go-jet/jet/v2/postgres"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
)

var validCourseNumberPattern = regexp2.MustCompile(`^\d{2}-\d{3}$`, 0)

type CourseIn struct {
	Number      string `json:"number"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Units       int32  `json:"units"`
	Description string `json:"description"`
}

func (c CourseIn) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Number, validation.Required, validation.By(MatchRegex(validCourseNumberPattern, "must be in the format ##-###"))),
		validation.Field(&c.Name, validation.Required, validation.Length(1, 0)),
		validation.Field(&c.Department, validation.Required, validation.Length(1, 0)),
		validation.Field(&c.Units, validation.Required, validation.Min(1)),
		validation.Field(&c.Description, validation.Required, validation.Length(1, 0)),
	)

}

type CourseIDIn struct {
	ID int64 `param:"id"`
}

func (c CourseIDIn) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required, validation.Min(1)),
	)
}

type CourseUpdateIn struct {
	CourseIDIn
	CourseIn
}

func (c CourseUpdateIn) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.CourseIDIn),
		validation.Field(&c.CourseIn),
	)
}

type CourseFilters struct {
	Number     *string `query:"number"`
	Name       *string `query:"name"`
	Department *string `query:"department"`
	Units      *int32  `query:"units"`
	PaginationIn
	SortIn
}

func (f CourseFilters) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Number, validation.NilOrNotEmpty, validation.Length(1, 0)),
		validation.Field(&f.Name, validation.NilOrNotEmpty, validation.Length(1, 0)),
		validation.Field(&f.Department, validation.NilOrNotEmpty, validation.Length(1, 0)),
		validation.Field(&f.Units, validation.NilOrNotEmpty, validation.Min(1)),
		validation.Field(&f.PaginationIn),
		validation.Field(&f.SortIn, validation.By(f.SortIn.ValidateNilOrNotEmptyIn("id", "number", "name", "-id", "-number", "-name"))),
	)
}

func (f CourseFilters) NewFilterCondition() pg.BoolExpression {
	condition := pg.Bool(true)
	if f.Number != nil {
		condition = condition.AND(IStringContains(tbl.Courses.Number, *f.Number))
	}
	if f.Name != nil {
		condition = condition.AND(IStringContains(tbl.Courses.Name, *f.Name))
	}
	if f.Department != nil {
		condition = condition.AND(IStringContains(tbl.Courses.Department, *f.Department))
	}
	if f.Units != nil {
		condition = condition.AND(tbl.Courses.Units.EQ(pg.Int32(*f.Units)))
	}
	return condition
}

var CourseSortMap = map[string]pg.Expression{
	"id":     tbl.Courses.ID,
	"number": tbl.Courses.Number,
	"name":   tbl.Courses.Name,
}

type CourseOut struct {
	ID         int64  `json:"id"`
	Number     string `json:"number"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Units      int32  `json:"units"`
}

func NewCourseModel(course *CourseIn) *model.Courses {
	if course == nil {
		course = &CourseIn{}
	}
	return &model.Courses{
		Number:      course.Number,
		Name:        course.Name,
		Department:  course.Department,
		Units:       course.Units,
		Description: course.Description,
	}
}

func NewCourseOut(course *model.Courses) *CourseOut {
	if course == nil {
		course = &model.Courses{}
	}
	return &CourseOut{
		ID:         course.ID,
		Number:     course.Number,
		Name:       course.Name,
		Department: course.Department,
		Units:      course.Units,
	}
}
