package schemas

import (
	pg "github.com/go-jet/jet/v2/postgres"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	tbl "github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"time"
)

type RatingIn struct {
	UserID     int64  `json:"-"`
	OfferingID int64  `json:"offering_id"`
	Overall    int32  `json:"overall"`
	Teaching   int32  `json:"teaching"`
	Materials  int32  `json:"materials"`
	Value      int32  `json:"value"`
	Difficulty int32  `json:"difficulty"`
	Workload   int32  `json:"workload"`
	Grading    int32  `json:"grading"`
	Comment    string `json:"comment"`
}

func (r RatingIn) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.OfferingID, validation.Required, validation.Min(1)),
		validation.Field(&r.Overall, validation.Required, validation.Min(1), validation.Max(5)),
		validation.Field(&r.Teaching, validation.Required, validation.Min(1), validation.Max(5)),
		validation.Field(&r.Materials, validation.Required, validation.Min(1), validation.Max(5)),
		validation.Field(&r.Value, validation.Required, validation.Min(1), validation.Max(5)),
		validation.Field(&r.Difficulty, validation.Required, validation.Min(1), validation.Max(5)),
		validation.Field(&r.Workload, validation.Required, validation.Min(1), validation.Max(5)),
		validation.Field(&r.Grading, validation.Required, validation.Min(1), validation.Max(5)),
		validation.Field(&r.Comment, validation.Required, validation.Length(1, 10000)),
	)
}

type RatingIDIn struct {
	ID int64 `param:"id"`
}

func (r RatingIDIn) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, validation.Min(1)),
	)
}

type RatingUpdateIn struct {
	RatingIDIn
	RatingIn
}

func (r RatingUpdateIn) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.RatingIDIn),
		validation.Field(&r.RatingIn),
	)
}

type RatingFilters struct {
	UserID        *int64  `query:"user_id"`
	CourseID      *int64  `query:"course_id"`
	Semester      *string `query:"semester"`
	Year          *int32  `query:"year"`
	InstructorIDs []int64 `query:"instructor_ids"`
	Overall       *int32  `query:"overall"`
	PaginationIn
	SortIn
}

func (f RatingFilters) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.UserID, validation.NilOrNotEmpty, validation.Min(1)),
		validation.Field(&f.CourseID, validation.NilOrNotEmpty, validation.Min(1)),
		validation.Field(&f.Semester, validation.NilOrNotEmpty, validation.In(SemesterList...)),
		validation.Field(&f.Year, validation.NilOrNotEmpty, validation.Min(1)),
		validation.Field(&f.InstructorIDs, validation.NilOrNotEmpty, validation.Each(validation.Min(1))),
		validation.Field(&f.Overall, validation.NilOrNotEmpty, validation.Min(1), validation.Max(5)),
		validation.Field(&f.PaginationIn),
		validation.Field(&f.SortIn, validation.By(f.SortIn.ValidateNilOrNotEmptyIn("id", "overall", "semester", "year", "created_at", "updated_at", "net_likes", "-id", "-overall", "-semester", "-year", "-created_at", "-updated_at", "-net_likes"))),
	)
}

func (f RatingFilters) NewFilterCondition() pg.BoolExpression {
	condition := pg.Bool(true)
	if f.UserID != nil {
		condition = condition.AND(tbl.Ratings.UserID.EQ(pg.Int64(*f.UserID)))
	}
	if f.CourseID != nil {
		condition = condition.AND(tbl.Offerings.CourseID.EQ(pg.Int64(*f.CourseID)))
	}
	if f.Semester != nil {
		condition = condition.AND(pg.LOWER(pg.CAST(tbl.Offerings.Semester).AS_TEXT()).EQ(pg.LOWER(pg.String(*f.Semester))))
	}
	if f.Year != nil {
		condition = condition.AND(tbl.Offerings.Year.EQ(pg.Int32(*f.Year)))
	}
	if f.InstructorIDs != nil {
		IDs := make([]pg.Expression, len(f.InstructorIDs))
		for i, ID := range f.InstructorIDs {
			IDs[i] = pg.Int64(ID)
		}
		condition = condition.AND(pg.EXISTS(
			tbl.Offerings.LEFT_JOIN(tbl.Teaches, tbl.Offerings.ID.EQ(tbl.Teaches.OfferingID)).
				SELECT(tbl.Teaches.InstructorID).
				WHERE(tbl.Offerings.ID.EQ(tbl.Ratings.OfferingID).
					AND(tbl.Teaches.InstructorID.IN(IDs...)))),
		)
	}
	if f.Overall != nil {
		condition = condition.AND(tbl.Ratings.Overall.EQ(pg.Int32(*f.Overall)))
	}
	return condition
}

var RatingSortMap = map[string]pg.Expression{
	"id":         tbl.Ratings.ID,
	"overall":    tbl.Ratings.Overall,
	"semester":   tbl.Offerings.Semester,
	"year":       tbl.Offerings.Year,
	"created_at": tbl.Ratings.CreatedAt,
	"updated_at": tbl.Ratings.UpdatedAt,
	"net_likes":  pg.IntegerColumn("net_likes"),
}

func NewRatingSortMapFrom(subQuery pg.SelectTable) map[string]pg.Expression {
	return map[string]pg.Expression{
		"id":         tbl.Ratings.ID.From(subQuery),
		"overall":    tbl.Ratings.Overall.From(subQuery),
		"semester":   tbl.Offerings.Semester.From(subQuery),
		"year":       tbl.Offerings.Year.From(subQuery),
		"created_at": tbl.Ratings.CreatedAt.From(subQuery),
		"updated_at": tbl.Ratings.UpdatedAt.From(subQuery),
		"net_likes":  pg.IntegerColumn("net_likes").From(subQuery),
	}
}

type RatingOut struct {
	ID               int64        `json:"id"`
	User             *UserOut     `json:"user,omitempty"`
	UserID           int64        `json:"user_id,omitempty"`
	Offering         *OfferingOut `json:"offering,omitempty"`
	OfferingID       int64        `json:"offering_id,,omitempty"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	Overall          int32        `json:"overall"`
	Teaching         int32        `json:"teaching"`
	Materials        int32        `json:"materials"`
	Value            int32        `json:"value"`
	Difficulty       int32        `json:"difficulty"`
	Workload         int32        `json:"workload"`
	Grading          int32        `json:"grading"`
	Comment          string       `json:"comment"`
	NetLikes         int64        `json:"net_likes"`
	LikedByViewer    bool         `json:"liked_by_viewer"`
	DislikedByViewer bool         `json:"disliked_by_viewer"`
}

type LikeOut struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RatingID  int64     `json:"rating_id"`
	CreatedAt time.Time `json:"created_at"`
}

type DislikeOut struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RatingID  int64     `json:"rating_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RatingStatsOut struct {
	RatingCount   int64         `json:"rating_count"`
	Overall       map[int]int64 `json:"overall"`
	Teaching      map[int]int64 `json:"teaching"`
	Materials     map[int]int64 `json:"materials"`
	Value         map[int]int64 `json:"value"`
	Difficulty    map[int]int64 `json:"difficulty"`
	Workload      map[int]int64 `json:"workload"`
	Grading       map[int]int64 `json:"grading"`
	AvgOverall    float64       `json:"avg_overall"`
	AvgTeaching   float64       `json:"avg_teaching"`
	AvgMaterials  float64       `json:"avg_materials"`
	AvgValue      float64       `json:"avg_value"`
	AvgDifficulty float64       `json:"avg_difficulty"`
	AvgWorkload   float64       `json:"avg_workload"`
	AvgGrading    float64       `json:"avg_grading"`
}

func NewRatingModel(rating *RatingIn) *model.Ratings {
	if rating == nil {
		rating = &RatingIn{}
	}
	return &model.Ratings{
		UserID:     rating.UserID,
		OfferingID: rating.OfferingID,
		UpdatedAt:  time.Now(),
		Overall:    rating.Overall,
		Teaching:   rating.Teaching,
		Materials:  rating.Materials,
		Value:      rating.Value,
		Difficulty: rating.Difficulty,
		Workload:   rating.Workload,
		Grading:    rating.Grading,
		Comment:    rating.Comment,
	}
}

func NewRatingOut(rating *model.Ratings, user *UserOut, offering *OfferingOut, netLikes int64, likedByViewer, dislikedByView bool) *RatingOut {
	if rating == nil {
		rating = &model.Ratings{}
	}
	if user == nil {
		user = &UserOut{}
	}
	if offering == nil {
		offering = &OfferingOut{}
	}
	return &RatingOut{
		ID:               rating.ID,
		User:             user,
		Offering:         offering,
		CreatedAt:        rating.CreatedAt,
		UpdatedAt:        rating.UpdatedAt,
		Overall:          rating.Overall,
		Teaching:         rating.Teaching,
		Materials:        rating.Materials,
		Value:            rating.Value,
		Difficulty:       rating.Difficulty,
		Workload:         rating.Workload,
		Grading:          rating.Grading,
		Comment:          rating.Comment,
		NetLikes:         netLikes,
		LikedByViewer:    likedByViewer,
		DislikedByViewer: dislikedByView,
	}
}

func NewRatingOutBrief(ratingIn *RatingIn, ratingModel *model.Ratings) *RatingOut {
	if ratingIn == nil {
		ratingIn = &RatingIn{}
	}
	if ratingModel == nil {
		ratingModel = &model.Ratings{}
	}
	return &RatingOut{
		ID:         ratingModel.ID,
		UserID:     ratingIn.UserID,
		OfferingID: ratingIn.OfferingID,
		CreatedAt:  ratingModel.CreatedAt,
		UpdatedAt:  ratingModel.UpdatedAt,
		Overall:    ratingIn.Overall,
		Teaching:   ratingIn.Teaching,
		Materials:  ratingIn.Materials,
		Value:      ratingIn.Value,
		Difficulty: ratingIn.Difficulty,
		Workload:   ratingIn.Workload,
		Grading:    ratingIn.Grading,
		Comment:    ratingIn.Comment,
	}
}

func FilterRatingsSQL(sortIn *SortIn, paginationIn *PaginationIn, condition pg.BoolExpression, viewerID int64) (pg.SelectStatement, error) {
	if sortIn == nil {
		sortIn = new(SortIn)
	}

	sort, err := sortIn.NewOrderByArrays(RatingSortMap)
	if err != nil {
		return nil, err
	}

	likesUserIDs := tbl.Likes.SELECT(tbl.Likes.UserID).
		WHERE(tbl.Likes.RatingID.EQ(tbl.Ratings.ID)).
		AsTable("likes_user_ids_table")
	dislikesUserIDs := tbl.Dislikes.SELECT(tbl.Dislikes.UserID).
		WHERE(tbl.Dislikes.RatingID.EQ(tbl.Ratings.ID)).
		AsTable("dislikes_user_ids_table")

	likesCount := pg.LATERAL(likesUserIDs.SELECT(pg.COUNT(tbl.Likes.UserID.From(likesUserIDs)).AS("count"))).
		AS("likes_count")
	dislikesCount := pg.LATERAL(dislikesUserIDs.SELECT(pg.COUNT(tbl.Dislikes.UserID.From(dislikesUserIDs)).AS("count"))).
		AS("dislikes_count")
	likedByViewer := pg.EXISTS(
		likesUserIDs.SELECT(tbl.Likes.UserID.From(likesUserIDs)).
			WHERE(tbl.Likes.UserID.From(likesUserIDs).EQ(pg.Int64(viewerID))),
	)
	dislikedByViewer := pg.EXISTS(
		dislikesUserIDs.SELECT(tbl.Dislikes.UserID.From(dislikesUserIDs)).
			WHERE(tbl.Dislikes.UserID.From(dislikesUserIDs).EQ(pg.Int64(viewerID))),
	)

	filteredRatings := tbl.Ratings.INNER_JOIN(tbl.Users, tbl.Ratings.UserID.EQ(tbl.Users.ID)).
		INNER_JOIN(tbl.Offerings, tbl.Offerings.ID.EQ(tbl.Ratings.OfferingID)).
		LEFT_JOIN(likesCount, pg.Bool(true)).
		LEFT_JOIN(dislikesCount, pg.Bool(true)).
		SELECT(
			pg.COUNT(tbl.Ratings.ID).OVER().AS("total_records"),
			tbl.Ratings.AllColumns,
			tbl.Users.AllColumns,
			tbl.Offerings.AllColumns,
			pg.IntegerColumn("count").From(likesCount).SUB(pg.IntegerColumn("count").From(dislikesCount)).
				AS("net_likes"),
			likedByViewer.AS("liked_by_viewer"),
			dislikedByViewer.AS("disliked_by_viewer"),
		).
		WHERE(condition).
		ORDER_BY(sort...)

	if paginationIn != nil {
		filteredRatings = filteredRatings.LIMIT(paginationIn.Limit()).OFFSET(paginationIn.Offset())
	}

	return filteredRatings, nil
}

func FilterAndAnnotateRatingsSQL(sortIn *SortIn, paginationIn *PaginationIn, condition pg.BoolExpression, viewerID int64) (pg.SelectStatement, error) {
	if sortIn == nil {
		sortIn = new(SortIn)
	}

	filteredRatings, err := FilterRatingsSQL(sortIn, paginationIn, condition, viewerID)
	if err != nil {
		return nil, err
	}
	filteredRatingsTable := filteredRatings.AsTable("filtered_ratings")

	sort, err := sortIn.NewOrderByArrays(NewRatingSortMapFrom(filteredRatingsTable))
	if err != nil {
		return nil, err
	}
	sort = append(sort, tbl.Instructors.ID.ASC())

	stmt := filteredRatingsTable.INNER_JOIN(tbl.Courses, tbl.Courses.ID.EQ(tbl.Offerings.CourseID.From(filteredRatingsTable))).
		LEFT_JOIN(tbl.Teaches, tbl.Offerings.ID.From(filteredRatingsTable).EQ(tbl.Teaches.OfferingID)).
		INNER_JOIN(tbl.Instructors, tbl.Teaches.InstructorID.EQ(tbl.Instructors.ID)).
		SELECT(filteredRatingsTable.AllColumns(), tbl.Courses.AllColumns, tbl.Instructors.AllColumns).
		ORDER_BY(sort...)

	return stmt, nil
}

func NewLikeModel(ratingID, userID int64) *model.Likes {
	return &model.Likes{
		UserID:   userID,
		RatingID: ratingID,
	}
}

func NewLikeOut(like *model.Likes) *LikeOut {
	if like == nil {
		like = &model.Likes{}
	}
	return &LikeOut{
		ID:        like.ID,
		UserID:    like.UserID,
		RatingID:  like.RatingID,
		CreatedAt: like.CreatedAt,
	}
}

func NewDislikeModel(ratingID, userID int64) *model.Dislikes {
	return &model.Dislikes{
		UserID:   userID,
		RatingID: ratingID,
	}
}

func NewDislikeOut(dislike *model.Dislikes) *DislikeOut {
	if dislike == nil {
		dislike = &model.Dislikes{}
	}
	return &DislikeOut{
		ID:        dislike.ID,
		UserID:    dislike.UserID,
		RatingID:  dislike.RatingID,
		CreatedAt: dislike.CreatedAt,
	}
}

type RatingStatsModel struct {
	RatingCount   int64
	AvgOverall    float64
	AvgTeaching   float64
	AvgMaterials  float64
	AvgValue      float64
	AvgDifficulty float64
	AvgWorkload   float64
	AvgGrading    float64
	Overall1      int64
	Overall2      int64
	Overall3      int64
	Overall4      int64
	Overall5      int64
	Teaching1     int64
	Teaching2     int64
	Teaching3     int64
	Teaching4     int64
	Teaching5     int64
	Materials1    int64
	Materials2    int64
	Materials3    int64
	Materials4    int64
	Materials5    int64
	Value1        int64
	Value2        int64
	Value3        int64
	Value4        int64
	Value5        int64
	Difficulty1   int64
	Difficulty2   int64
	Difficulty3   int64
	Difficulty4   int64
	Difficulty5   int64
	Workload1     int64
	Workload2     int64
	Workload3     int64
	Workload4     int64
	Workload5     int64
	Grading1      int64
	Grading2      int64
	Grading3      int64
	Grading4      int64
	Grading5      int64
}

func NewRatingStatsOut(stats *RatingStatsModel) *RatingStatsOut {
	return &RatingStatsOut{
		RatingCount:   stats.RatingCount,
		Overall:       map[int]int64{1: stats.Overall1, 2: stats.Overall2, 3: stats.Overall3, 4: stats.Overall4, 5: stats.Overall5},
		Teaching:      map[int]int64{1: stats.Teaching1, 2: stats.Teaching2, 3: stats.Teaching3, 4: stats.Teaching4, 5: stats.Teaching5},
		Materials:     map[int]int64{1: stats.Materials1, 2: stats.Materials2, 3: stats.Materials3, 4: stats.Materials4, 5: stats.Materials5},
		Value:         map[int]int64{1: stats.Value1, 2: stats.Value2, 3: stats.Value3, 4: stats.Value4, 5: stats.Value5},
		Difficulty:    map[int]int64{1: stats.Difficulty1, 2: stats.Difficulty2, 3: stats.Difficulty3, 4: stats.Difficulty4, 5: stats.Difficulty5},
		Workload:      map[int]int64{1: stats.Workload1, 2: stats.Workload2, 3: stats.Workload3, 4: stats.Workload4, 5: stats.Workload5},
		Grading:       map[int]int64{1: stats.Grading1, 2: stats.Grading2, 3: stats.Grading3, 4: stats.Grading4, 5: stats.Grading5},
		AvgOverall:    stats.AvgOverall,
		AvgTeaching:   stats.AvgTeaching,
		AvgMaterials:  stats.AvgMaterials,
		AvgValue:      stats.AvgValue,
		AvgDifficulty: stats.AvgDifficulty,
		AvgWorkload:   stats.AvgWorkload,
		AvgGrading:    stats.AvgGrading,
	}
}
