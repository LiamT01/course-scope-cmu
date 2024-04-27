package schemas

import (
	"errors"
	"github.com/dlclark/regexp2"
	"github.com/go-jet/jet/v2/postgres"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/table"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var validUsernamePattern = regexp2.MustCompile(`^[\w.@+-]{1,30}$`, 0)

var validPasswordPattern = regexp2.MustCompile(`((?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[\W]).{8,64})`, 0)

var validAndrewIDPattern = regexp2.MustCompile(`^(?=.{1,20}$)[a-z]+[0-9]*$`, 0)

var AnonymousUser = &model.Users{}

func IsAnonymousUser(user *model.Users) bool {
	return user == AnonymousUser
}

func IsActivatedUser(user *model.Users) bool {
	return !IsAnonymousUser(user) && user != nil && user.Activated
}

func IsAdminUser(user *model.Users) bool {
	return !IsAnonymousUser(user) && user != nil && user.Admin
}

type AndrewIDIn struct {
	AndrewID string `json:"andrew_id"`
}

func (a AndrewIDIn) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AndrewID, validation.Required, validation.By(MatchRegex(validAndrewIDPattern, "must start with lowercase letters and optionally end with numbers, and be 1-20 characters long"))),
	)
}

type UserIn struct {
	Username string `json:"username"`
	AndrewIDIn
	PasswordIn
}

func (u UserIn) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.By(MatchRegex(validUsernamePattern, "must be 1-30 characters. Letters, digits and @/./+/-/_ only"))),
		validation.Field(&u.AndrewID, validation.Required, validation.By(MatchRegex(validAndrewIDPattern, "must start with lowercase letters and optionally end with numbers, and be 1-20 characters long"))),
		validation.Field(&u.PasswordIn),
	)
}

type UserIDIn struct {
	ID int64 `param:"id"`
}

func (u UserIDIn) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required, validation.Min(1)),
	)
}

type UserUpdateIn struct {
	UserIDIn
	UserIn
}

func (u UserUpdateIn) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserIDIn),
		validation.Field(&u.UserIn),
	)
}

type PasswordIn struct {
	Plaintext string `json:"password"`
}

func (p PasswordIn) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Plaintext, validation.Required, validation.By(MatchRegex(validPasswordPattern, "must contain at least one digit, one lowercase letter, one uppercase letter, one special character, and be 8-64 characters long"))),
	)
}

func HashPasswordPlaintext(plaintext string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func CheckPassword(plaintext string, hash []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(plaintext)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

type UserFilters struct {
	PaginationIn
	SortIn
}

func (f UserFilters) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.PaginationIn),
		validation.Field(&f.SortIn, validation.By(f.SortIn.ValidateNilOrNotEmptyIn("id", "username", "-id", "-username"))),
	)
}

var UserSortMap = map[string]postgres.Expression{
	"id":       table.Users.ID,
	"username": table.Users.Username,
}

type UsernameIn struct {
	Username string `json:"username"`
}

func (u UsernameIn) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.By(MatchRegex(validUsernamePattern, "must be 1-30 characters. Letters, digits and @/./+/-/_ only"))),
	)
}

type UserOut struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Activated bool      `json:"activated"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStatsOut struct {
	LikesReceived int64 `json:"likes_received"`
	CoursesRated  int64 `json:"courses_rated"`
}

func NewUserModel(user *UserIn) (*model.Users, error) {
	if user == nil {
		user = &UserIn{}
	}
	hash, err := HashPasswordPlaintext(user.PasswordIn.Plaintext)
	if err != nil {
		return nil, err
	}

	return &model.Users{
		AndrewID:     user.AndrewID,
		Username:     user.Username,
		PasswordHash: hash,
	}, nil
}

func NewUserOut(user *model.Users) *UserOut {
	if user == nil {
		user = &model.Users{}
	}
	return &UserOut{
		ID:        user.ID,
		Username:  user.Username,
		Activated: user.Activated,
		CreatedAt: user.CreatedAt,
	}
}

func NewUserStatsOut(likesReceived, coursesRated int64) *UserStatsOut {
	return &UserStatsOut{
		LikesReceived: likesReceived,
		CoursesRated:  coursesRated,
	}
}
