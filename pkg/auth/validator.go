package auth

import "errors"

var (
	ErrUsernameTooLong  = errors.New("account: username too long")
	ErrUsernameTooShort = errors.New("account: username too short")
	ErrUsernameEmpty    = errors.New("account: username empty")
)

const minTitle = 5
const maxTitle = 32

func validateTitle(title string) error {
	l := len(title)

	switch {
	case l == 0:
		return ErrUsernameEmpty
	case l < minTitle:
		return ErrUsernameTooShort
	case l > maxTitle:
		return ErrUsernameTooLong
	default:
		return nil
	}
}
