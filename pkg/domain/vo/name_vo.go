package vo

import "github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"

var (
	ErrInvalidName = errors.NewBusinessError("invalid name")
)

type Name struct {
	value string
}

func (vo Name) Value() string {
	return vo.value
}

func NewNameVo(value string) (*Name, error) {
	if len(value) < 3 {
		return nil, ErrInvalidName
	}

	return &Name{value: value}, nil
}
