package flow

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// OriginPathTestSuite define a estrutura para o suite de testes de OriginPath
type OriginPathTestSuite struct {
	suite.Suite
}

// SetupTest configura antes de cada teste
func (suite *OriginPathTestSuite) SetupTest() {
	// Adicione aqui as configurações necessárias antes de cada teste
}

// TestNewOriginPathVoValid verifica a criação com um valor válido
func (suite *OriginPathTestSuite) TestNewOriginPathVoValid() {
	pathValue := "valid/path"
	originPath, err := NewOriginPathVo(pathValue)

	assert.Nil(suite.T(), err, "erro deve ser nil para um valor válido")
	assert.NotNil(suite.T(), originPath, "originPath não deve ser nil para um valor válido")
	assert.Equal(suite.T(), pathValue, originPath.Value(), "o valor de originPath deve ser igual ao valor fornecido")
}

// TestNewOriginPathVoInvalid verifica a criação com um valor inválido (string vazia)
func (suite *OriginPathTestSuite) TestNewOriginPathVoInvalid() {
	pathValue := ""
	originPath, err := NewOriginPathVo(pathValue)

	assert.NotNil(suite.T(), err, "erro deve ser retornado para um valor inválido")
	assert.Nil(suite.T(), originPath, "originPath deve ser nil para um valor inválido")
	assert.Equal(suite.T(), ErrInvalidOriginPath, err, "erro deve ser ErrInvalidOriginPath para um valor inválido")
}

// TestValue verifica o método Value para OriginPath
func (suite *OriginPathTestSuite) TestValue() {
	pathValue := "another/path"
	originPath := OriginPath{value: pathValue}

	assert.Equal(suite.T(), pathValue, originPath.Value(), "o valor do método Value deve ser igual ao valor fornecido")
}

// RunOriginPathTestSuite inicia o suite de testes
func TestOriginPathTestSuite(t *testing.T) {
	suite.Run(t, new(OriginPathTestSuite))
}
