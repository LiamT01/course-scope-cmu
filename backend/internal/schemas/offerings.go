package schemas

import (
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
)

var SemesterList = []interface{}{"Fall", "Spring", "Summer 1", "Summer 2", "Winter", "fall", "spring", "summer 1", "summer 2", "winter"}

type OfferingIn struct {
	CourseID      int64   `json:"course_id"`
	Semester      string  `json:"semester"`
	Year          int32   `json:"year"`
	Location      string  `json:"location"`
	InstructorIDs []int64 `json:"instructor_ids"`
}

func (o OfferingIn) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.CourseID, validation.Required, validation.Min(1)),
		validation.Field(&o.Semester, validation.Required, validation.In(SemesterList...)),
		validation.Field(&o.Year, validation.Required, validation.Min(1000), validation.Max(9999)),
		validation.Field(&o.Location, validation.Required, validation.Length(1, 150)),
		validation.Field(&o.InstructorIDs, validation.Required, validation.Length(1, 0), validation.Each(validation.Min(1))),
	)
}

type OfferingIDIn struct {
	ID int64 `param:"id"`
}

func (o OfferingIDIn) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.ID, validation.Required, validation.Min(1)),
	)
}

type OfferingUpdateIn struct {
	OfferingIDIn
	OfferingIn
}

func (o OfferingUpdateIn) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.OfferingIDIn),
		validation.Field(&o.OfferingIn),
	)
}

type OfferingFilters struct {
	CourseID      *int64  `query:"course_id"`
	Semester      *string `query:"semester"`
	Year          *int32  `query:"year"`
	Location      *string `query:"location"`
	InstructorIDs []int64 `query:"instructor_ids"`
	PaginationIn
	SortIn
}

func (f OfferingFilters) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.CourseID, validation.NilOrNotEmpty, validation.Min(1)),
		validation.Field(&f.Semester, validation.NilOrNotEmpty, validation.In(SemesterList...)),
		validation.Field(&f.Year, validation.NilOrNotEmpty, validation.Min(1000), validation.Max(9999)),
		validation.Field(&f.Location, validation.NilOrNotEmpty, validation.Length(1, 150)),
		validation.Field(&f.InstructorIDs, validation.NilOrNotEmpty, validation.Each(validation.Min(1))),
		validation.Field(&f.PaginationIn),
		validation.Field(&f.SortIn, validation.By(f.SortIn.ValidateNilOrNotEmptyIn("id", "semester", "year", "-id", "-semester", "-year"))),
	)
}

func (f OfferingFilters) NewFilterCondition() pg.BoolExpression {
	condition := pg.Bool(true)
	if f.CourseID != nil {
		condition = condition.AND(tbl.Offerings.CourseID.EQ(pg.Int64(*f.CourseID)))
	}
	if f.Semester != nil {
		condition = condition.AND(pg.LOWER(pg.CAST(tbl.Offerings.Semester).AS_TEXT()).EQ(pg.LOWER(pg.String(*f.Semester))))
	}
	if f.Year != nil {
		condition = condition.AND(tbl.Offerings.Year.EQ(pg.Int32(*f.Year)))
	}
	if f.Location != nil {
		condition = condition.AND(IStringContains(tbl.Offerings.Location, *f.Location))
	}
	if f.InstructorIDs != nil {
		IDs := make([]pg.Expression, len(f.InstructorIDs))
		for i, id := range f.InstructorIDs {
			IDs[i] = pg.Int64(id)
		}
		condition = condition.AND(
			pg.EXISTS(
				tbl.Teaches.SELECT(tbl.Teaches.InstructorID).
					WHERE(tbl.Teaches.OfferingID.EQ(tbl.Offerings.ID).
						AND(tbl.Teaches.InstructorID.IN(IDs...))),
			),
		)
	}
	return condition
}

var OfferingSortMap = map[string]pg.Expression{
	"id":       tbl.Offerings.ID,
	"semester": tbl.Offerings.Semester,
	"year":     tbl.Offerings.Year,
}

func NewOfferingSortMapFrom(subQuery pg.SelectTable) map[string]pg.Expression {
	return map[string]pg.Expression{
		"id":       tbl.Offerings.ID.From(subQuery),
		"semester": tbl.Offerings.Semester.From(subQuery),
		"year":     tbl.Offerings.Year.From(subQuery),
	}
}

type OfferingOut struct {
	ID            int64            `json:"id"`
	CourseID      int64            `json:"course_id,omitempty"`
	Course        *CourseOut       `json:"course,omitempty"`
	Semester      string           `json:"semester"`
	Year          int32            `json:"year"`
	Location      string           `json:"location"`
	Instructors   []*InstructorOut `json:"instructors,omitempty"`
	InstructorIDs []int64          `json:"instructor_ids,omitempty"`
}

func NewOfferingModel(offering *OfferingIn) *model.Offerings {
	if offering == nil {
		offering = &OfferingIn{}
	}
	return &model.Offerings{
		CourseID: offering.CourseID,
		Semester: model.SemesterType(offering.Semester),
		Year:     offering.Year,
		Location: offering.Location,
	}
}

func NewTeachesModel(offeringID int64, instructorID int64) *model.Teaches {
	return &model.Teaches{
		OfferingID:   offeringID,
		InstructorID: instructorID,
	}
}

func NewOfferingOut(offering *model.Offerings, course *CourseOut, instructors []*InstructorOut) *OfferingOut {
	if offering == nil {
		offering = &model.Offerings{}
	}
	if course == nil {
		course = &CourseOut{}
	}
	if instructors == nil {
		instructors = []*InstructorOut{}
	}
	return &OfferingOut{
		ID:          offering.ID,
		Course:      course,
		Semester:    string(offering.Semester),
		Year:        offering.Year,
		Location:    offering.Location,
		Instructors: instructors,
	}
}

func NewOfferingOutBrief(offeringIn *OfferingIn, offeringModel *model.Offerings) *OfferingOut {
	if offeringIn == nil {
		offeringIn = &OfferingIn{}
	}
	if offeringModel == nil {
		offeringModel = &model.Offerings{}
	}
	return &OfferingOut{
		ID:            offeringModel.ID,
		CourseID:      offeringIn.CourseID,
		Semester:      string(offeringModel.Semester),
		Year:          offeringModel.Year,
		Location:      offeringModel.Location,
		InstructorIDs: offeringIn.InstructorIDs,
	}
}
