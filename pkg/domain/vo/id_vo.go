package vo

import (
	"fmt"
	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

// Value retorna o valor UUID de ID.
func (id ID) Value() uuid.UUID {
	return id.value
}

// NewID cria uma nova instância de ID com base no valor fornecido.
func NewID(value interface{}) (*ID, error) {
	switch v := value.(type) {
	case string:
		id, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		return &ID{value: id}, nil
	case uuid.UUID:
		return &ID{value: v}, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}
