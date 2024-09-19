package vo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// NetworkAddressTestSuite estrutura para os testes de NetworkAddress
type NetworkAddressTestSuite struct {
	suite.Suite
	ctx context.Context
}

// SetupTest inicializa o contexto para os testes
func (suite *NetworkAddressTestSuite) SetupTest() {
	suite.ctx = context.Background()
}

// TestValidIP verifica se um IP válido retorna sem erros
func (suite *NetworkAddressTestSuite) TestValidIP() {
	ip := "192.168.1.1"
	na, err := NewNetworkAddress(suite.ctx, ip)
	assert.Nil(suite.T(), err, "erro deve ser nil para um IP válido")
	assert.Equal(suite.T(), ip, na.value, "o valor do IP deve ser o mesmo que o passado")
}

// TestInvalidIP verifica se um IP inválido retorna um erro
func (suite *NetworkAddressTestSuite) TestInvalidIP() {
	ip := "999.999.999.999"
	na, err := NewNetworkAddress(suite.ctx, ip)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para um IP inválido")
	assert.Nil(suite.T(), na, "NetworkAddress deve ser nil para um IP inválido")
}

// TestValidHostname verifica se um hostname válido resolve corretamente
func (suite *NetworkAddressTestSuite) TestValidHostname() {
	hostname := "localhost"
	na, err := NewNetworkAddress(suite.ctx, hostname)
	assert.Nil(suite.T(), err, "erro deve ser nil para um hostname válido")
	assert.Equal(suite.T(), hostname, na.value, "o valor do hostname deve ser o mesmo que o passado")
}

// TestInvalidHostname verifica se um hostname inválido retorna um erro
func (suite *NetworkAddressTestSuite) TestInvalidHostname() {
	hostname := "invalidhostname"
	na, err := NewNetworkAddress(suite.ctx, hostname)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para um hostname inválido")
	assert.Nil(suite.T(), na, "NetworkAddress deve ser nil para um hostname inválido")
}

// TestNoIPAssociatedWithHost verifica se um hostname sem IPs associados retorna um erro
func (suite *NetworkAddressTestSuite) TestNoIPAssociatedWithHost() {
	hostname := "noip.example.com"
	na, err := NewNetworkAddress(suite.ctx, hostname)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para um hostname sem IPs associados")
	assert.Nil(suite.T(), na, "NetworkAddress deve ser nil para um hostname sem IPs associados")
}

// TestNetworkAddressTestSuite inicia o suite de teste
func TestNetworkAddressTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkAddressTestSuite))
}
