package vo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// NetworkPortTestSuite estrutura para os testes de NetworkPort
type NetworkPortTestSuite struct {
	suite.Suite
}

// TestValidPortInt verifica se um valor de porta inteiro válido é tratado corretamente
func (suite *NetworkPortTestSuite) TestValidPortInt() {
	port := 8080
	np, err := NewNetworkPortVo(port)
	assert.Nil(suite.T(), err, "erro deve ser nil para um valor de porta inteiro válido")
	assert.Equal(suite.T(), port, np.Value(), "o valor da porta deve ser igual ao valor fornecido")
}

// TestValidPortString verifica se um valor de porta em string válido é tratado corretamente
func (suite *NetworkPortTestSuite) TestValidPortString() {
	port := "8080"
	expectedValue := 8080
	np, err := NewNetworkPortVo(port)
	assert.Nil(suite.T(), err, "erro deve ser nil para um valor de porta em string válido")
	assert.Equal(suite.T(), expectedValue, np.Value(), "o valor da porta deve ser igual ao valor fornecido")
}

// TestInvalidPortString verifica se um valor de porta em string inválido retorna erro
func (suite *NetworkPortTestSuite) TestInvalidPortString() {
	port := "invalid"
	np, err := NewNetworkPortVo(port)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para um valor de porta em string inválido")
	assert.Nil(suite.T(), np, "NetworkPort deve ser nil para um valor de porta em string inválido")
}

// TestInvalidPortRange verifica se um valor de porta fora do intervalo permitido retorna erro
func (suite *NetworkPortTestSuite) TestInvalidPortRange() {
	port := 70000
	np, err := NewNetworkPortVo(port)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para um valor de porta fora do intervalo permitido")
	assert.Nil(suite.T(), np, "NetworkPort deve ser nil para um valor de porta fora do intervalo permitido")
}

// TestPortNotNumber verifica se um valor de porta que não é número retorna erro
func (suite *NetworkPortTestSuite) TestPortNotNumber() {
	port := []byte{1, 2, 3}
	np, err := NewNetworkPortVo(port)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para um valor de porta que não é número")
	assert.Nil(suite.T(), np, "NetworkPort deve ser nil para um valor de porta que não é número")
}

// TestNetworkPortTestSuite inicia o suite de teste
func TestNetworkPortTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkPortTestSuite))
}
