package schemas

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	"time"
)

var (
	ScopeList     = []interface{}{model.ScopeType_Act.String(), model.ScopeType_Auth.String(), model.ScopeType_Pwd.String()}
	SentScopeList = []interface{}{model.ScopeType_Act.String(), model.ScopeType_Pwd.String()}
)

type TokenIn struct {
	Token string `json:"token" query:"token"`
}

func (t TokenIn) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Token, validation.Required, validation.Length(26, 26)),
	)
}

type TokenOut struct {
	Token  string          `json:"token"`
	UserID int64           `json:"-"`
	Hash   []byte          `json:"-"`
	Expiry time.Time       `json:"expiry"`
	TTL    time.Duration   `json:"-"`
	Scope  model.ScopeType `json:"scope"`
}

type CredentialsIn struct {
	AndrewIDIn
	PasswordIn
}

func (c CredentialsIn) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.AndrewIDIn),
		validation.Field(&c.PasswordIn),
	)
}

type ResetPasswordIn struct {
	TokenIn
	PasswordIn
}

func (r ResetPasswordIn) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.TokenIn),
		validation.Field(&r.PasswordIn),
	)
}

func HashTokenPlaintext(plaintext string) []byte {
	hashArray := sha256.Sum256([]byte(plaintext))
	return hashArray[:]
}

func NewTokenOut(userID int64, ttl time.Duration, scope model.ScopeType) (*TokenOut, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return nil, err
	}

	plaintext := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := HashTokenPlaintext(plaintext)

	return &TokenOut{
		Token:  plaintext,
		UserID: userID,
		Hash:   hash,
		Expiry: time.Now().Add(ttl),
		TTL:    ttl,
		Scope:  scope,
	}, nil
}

func NewTokenModel(t *TokenOut) *model.Tokens {
	if t == nil {
		t = &TokenOut{}
	}
	return &model.Tokens{
		Hash:   t.Hash,
		UserID: t.UserID,
		Expiry: t.Expiry,
		Scope:  t.Scope,
	}
}
