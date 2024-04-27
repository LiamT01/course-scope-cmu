package schemas

import (
	"errors"
	"github.com/dlclark/regexp2"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-ozzo/ozzo-validation"
	"strings"
)

func MatchRegex(pattern *regexp2.Regexp, errorMsg string) validation.RuleFunc {
	return func(v interface{}) error {
		s := v.(string)
		match, _ := pattern.MatchString(s)
		if !match {
			return errors.New(errorMsg)
		}
		return nil
	}
}

func IStringContains(column postgres.ColumnString, s string) postgres.BoolExpression {
	return postgres.LOWER(column).LIKE(postgres.String("%" + strings.ToLower(s) + "%"))
}
