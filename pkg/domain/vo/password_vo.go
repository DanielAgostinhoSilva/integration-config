package vo

import (
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"unicode"
)

var (
	ErrPasswordIsTooShort         = errors.NewBusinessError("password is too short, must be at least 8 characters long")
	ErrPasswordUppercaseLetter    = errors.NewBusinessError("password must include at least one uppercase letter")
	ErrPasswordLowercaseLetter    = errors.NewBusinessError("password must include at least one lowercase letter")
	ErrPasswordAtLeastOneDigit    = errors.NewBusinessError("password must include at least one digit")
	ErrPasswordAtLeastSpecialChar = errors.NewBusinessError("password must include at least one special character")
)

type PasswordVo struct {
	value string
}

func (p PasswordVo) Value() string {
	return p.value
}

func NewPasswordVo(password string) (*PasswordVo, error) {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	if len(password) < 8 {
		return nil, ErrPasswordIsTooShort
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		}
	}

	if !upp {
		return nil, ErrPasswordLowercaseLetter
	}
	if !low {
		return nil, ErrPasswordUppercaseLetter
	}
	if !num {
		return nil, ErrPasswordAtLeastOneDigit
	}
	if !sym {
		return nil, ErrPasswordAtLeastSpecialChar
	}

	return &PasswordVo{value: password}, nil
}
