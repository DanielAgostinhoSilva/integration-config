package flow

import "github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"

var (
	ErrInvalidTargetPath = errors.NewBusinessError("invalid targetPath")
)

type TargetPath struct {
	value string
}

func (vo TargetPath) Value() string {
	return vo.value
}

func NewTargetPathVo(value string) (*TargetPath, error) {
	if len(value) == 0 {
		return nil, ErrInvalidTargetPath
	}

	return &TargetPath{value: value}, nil
}
