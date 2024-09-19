package vo

import (
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"strconv"
)

var (
	ErrCouldNotConvertPort = errors.NewBusinessError("could not convert port to integer")
	ErrInvalidPort         = errors.NewBusinessError("invalid port, port is out of the allowed range (0-65535)")
	ErrPortNotNumber       = errors.NewBusinessError("port is not a number")
)

// NetworkPort estrutura que representa uma porta de rede
type NetworkPort struct {
	value int
}

// Value retorna o valor da porta de rede
func (n NetworkPort) Value() int {
	return n.value
}

// NewNetworkPortVo cria uma nova instância de NetworkPort após validação
func NewNetworkPortVo(value any) (*NetworkPort, error) {
	port, err := parsePort(value)
	if err != nil {
		return nil, err
	}

	if !isValidPort(port) {
		return nil, ErrInvalidPort
	}

	return &NetworkPort{value: port}, nil
}

// parsePort tenta converter o valor fornecido para um inteiro representando a porta
func parsePort(value any) (int, error) {
	switch p := value.(type) {
	case string:
		portInt, err := strconv.Atoi(p)
		if err != nil {
			return 0, ErrCouldNotConvertPort
		}
		return portInt, nil
	case int:
		return p, nil
	default:
		return 0, ErrPortNotNumber
	}
}

// isValidPort verifica se a porta está no intervalo permitido (0-65535)
func isValidPort(port int) bool {
	return port >= 0 && port <= 65535
}
