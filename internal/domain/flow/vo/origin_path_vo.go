package flow

import "github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"

var (
	ErrInvalidOriginPath = errors.NewBusinessError("invalid originPath")
)

type OriginPath struct {
	value string
}

func (vo OriginPath) Value() string {
	return vo.value
}

func NewOriginPathVo(value string) (*OriginPath, error) {
	if len(value) == 0 {
		return nil, ErrInvalidOriginPath
	}

	return &OriginPath{value: value}, nil
}
