package schemas

import (
	"github.com/dlclark/regexp2"
	pg "github.com/go-jet/jet/v2/postgres"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
)

var validInstructorNamePattern = regexp2.MustCompile(`^[A-Za-z-']+,\s[A-Za-z-']+$`, 0)

type InstructorIn struct {
	Name string `json:"name"`
}

func (i InstructorIn) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Name, validation.Required, validation.By(MatchRegex(validInstructorNamePattern, "must be in the format 'Last, First'"))),
	)
}

type InstructorIDIn struct {
	ID int64 `param:"id"`
}

func (i InstructorIDIn) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.ID, validation.Required, validation.Min(1)),
	)
}

type InstructorUpdateIn struct {
	InstructorIDIn
	InstructorIn
}

func (i InstructorUpdateIn) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.InstructorIDIn),
		validation.Field(&i.InstructorIn),
	)
}

type InstructorFilters struct {
	Name *string `query:"name"`
	PaginationIn
	SortIn
}

func (f InstructorFilters) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.NilOrNotEmpty),
		validation.Field(&f.PaginationIn),
		validation.Field(&f.SortIn, validation.By(f.SortIn.ValidateNilOrNotEmptyIn("id", "name", "-id", "-name"))),
	)
}

func (f InstructorFilters) NewFilterCondition() pg.BoolExpression {
	condition := pg.Bool(true)
	if f.Name != nil {
		condition = condition.AND(IStringContains(tbl.Instructors.Name, *f.Name))
	}
	return condition
}

var InstructorSortMap = map[string]pg.Expression{
	"id":   tbl.Instructors.ID,
	"name": tbl.Instructors.Name,
}

type InstructorOut struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewInstructorModel(instructor *InstructorIn) *model.Instructors {
	if instructor == nil {
		instructor = &InstructorIn{}
	}
	return &model.Instructors{
		Name: instructor.Name,
	}
}

func NewInstructorOut(instructor *model.Instructors) *InstructorOut {
	if instructor == nil {
		instructor = &model.Instructors{}
	}
	return &InstructorOut{
		ID:   instructor.ID,
		Name: instructor.Name,
	}
}

func NewInstructorOurArray(instructors []*model.Instructors) []*InstructorOut {
	instructorOuts := make([]*InstructorOut, len(instructors))
	for i, instructor := range instructors {
		instructorOuts[i] = NewInstructorOut(instructor)
	}
	return instructorOuts
}
