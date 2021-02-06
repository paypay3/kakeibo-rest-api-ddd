package vo

import (
	"regexp"

	"github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/errors"
)

type Email string

const (
	maxEmailLength = 256
	emailFormat    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
)

var emailRegex = regexp.MustCompile(emailFormat)

func NewEmail(email string) (Email, error) {
	if len(email) == 0 ||
		len(email) > maxEmailLength ||
		!emailRegex.MatchString(email) {
		return "", errors.ErrInvalidEmail
	}

	return Email(email), nil
}

func (e Email) Value() string {
	return string(e)
}
