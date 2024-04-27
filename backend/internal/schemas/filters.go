package schemas

import (
	"errors"
	pg "github.com/go-jet/jet/v2/postgres"
	validation "github.com/go-ozzo/ozzo-validation"
	"math"
	"strings"
)

type PaginationIn struct {
	Page     *int64 `query:"page"`
	PageSize *int64 `query:"page_size"`
}

func (p PaginationIn) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Page, validation.NilOrNotEmpty, validation.Min(1)),
		validation.Field(&p.PageSize, validation.NilOrNotEmpty, validation.Min(1), validation.Max(100)),
	)
}

func (p PaginationIn) Limit() int64 {
	result := int64(10)
	if p.PageSize != nil {
		result = *p.PageSize
	}
	return result
}

func (p PaginationIn) Offset() int64 {
	result := int64(0)
	if p.Page != nil {
		result = (*p.Page - 1) * p.Limit()
	}
	return result
}

func (p PaginationIn) PageNumber() int64 {
	result := int64(1)
	if p.Page != nil {
		result = *p.Page
	}
	return result
}

type SortIn struct {
	Sort *string `query:"sort"`
}

func (s SortIn) ValidateNilOrNotEmptyIn(safeList ...interface{}) validation.RuleFunc {
	return func(v interface{}) error {
		return validation.Validate(s.Sort, validation.NilOrNotEmpty, validation.In(safeList...))
	}
}

type SortTuple struct {
	Column string
	Asc    bool
}

func (s SortIn) NewSortTuple() (result *SortTuple, ok bool) {
	if s.Sort == nil {
		return nil, false
	}
	desc := strings.HasPrefix(*s.Sort, "-")
	return &SortTuple{
		Column: strings.TrimPrefix(*s.Sort, "-"),
		Asc:    !desc,
	}, true
}

type Metadata struct {
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	FirstPage   int64 `json:"first_page"`
	LastPage    int64 `json:"last_page"`
	Total       int64 `json:"total"`
}

func NewMetadata(totalRecords int64, p *PaginationIn) *Metadata {
	return &Metadata{
		CurrentPage: p.PageNumber(),
		PageSize:    p.Limit(),
		FirstPage:   1,
		LastPage:    int64(math.Ceil(float64(totalRecords) / float64(p.Limit()))),
		Total:       totalRecords,
	}
}

//type Envelope map[string]interface{}

type Envelope struct {
	Items    interface{} `json:"items"`
	Metadata *Metadata   `json:"metadata"`
}

func NewEnvelope(items interface{}, totalRecords int64, p *PaginationIn) *Envelope {
	return &Envelope{
		Items:    items,
		Metadata: NewMetadata(totalRecords, p),
	}
}

func (s SortIn) NewOrderByArrays(sortMap map[string]pg.Expression) ([]pg.OrderByClause, error) {
	sort := make([]pg.OrderByClause, 0)
	sortColumn, ok := s.NewSortTuple()
	if ok {
		v, ok := sortMap[sortColumn.Column]
		if !ok || v == nil {
			return nil, errors.New("sort column not in sort map")
		}
		if sortColumn.Asc {
			sort = append(sort, v.ASC())
		} else {
			sort = append(sort, v.DESC())
		}
	}
	updateAt, ok := sortMap["updated_at"]
	if ok && updateAt != nil {
		sort = append(sort, updateAt.DESC())
	}
	ID, ok := sortMap["id"]
	if !ok || ID == nil {
		return nil, errors.New("id column not in sort map")
	}
	sort = append(sort, ID.ASC())
	return sort, nil
}
