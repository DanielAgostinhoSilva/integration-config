package flow

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TargetPathTestSuite define a estrutura para o suite de testes de TargetPath
type TargetPathTestSuite struct {
	suite.Suite
}

// SetupTest configura antes de cada teste
func (suite *TargetPathTestSuite) SetupTest() {
	// Adicione aqui as configurações necessárias antes de cada teste
}

// TestNewTargetPathVoValid verifica a criação com um valor válido
func (suite *TargetPathTestSuite) TestNewTargetPathVoValid() {
	pathValue := "valid/path"
	targetPath, err := NewTargetPathVo(pathValue)

	assert.Nil(suite.T(), err, "erro deve ser nil para um valor válido")
	assert.NotNil(suite.T(), targetPath, "targetPath não deve ser nil para um valor válido")
	assert.Equal(suite.T(), pathValue, targetPath.Value(), "o valor de targetPath deve ser igual ao valor fornecido")
}

// TestNewTargetPathVoInvalid verifica a criação com um valor inválido (string vazia)
func (suite *TargetPathTestSuite) TestNewTargetPathVoInvalid() {
	pathValue := ""
	targetPath, err := NewTargetPathVo(pathValue)

	assert.NotNil(suite.T(), err, "erro deve ser retornado para um valor inválido")
	assert.Nil(suite.T(), targetPath, "targetPath deve ser nil para um valor inválido")
	assert.Equal(suite.T(), ErrInvalidTargetPath, err, "erro deve ser ErrInvalidTargetPath para um valor inválido")
}

// TestValue verifica o método Value para TargetPath
func (suite *TargetPathTestSuite) TestValue() {
	pathValue := "another/path"
	targetPath := TargetPath{value: pathValue}

	assert.Equal(suite.T(), pathValue, targetPath.Value(), "o valor do método Value deve ser igual ao valor fornecido")
}

// TestTargetPathTestSuite inicia o suite de testes
func TestTargetPathTestSuite(t *testing.T) {
	suite.Run(t, new(TargetPathTestSuite))
}
