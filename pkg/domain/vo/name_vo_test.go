package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// NameTestSuite estrutura para os testes de Name
type NameTestSuite struct {
	suite.Suite
}

// TestNewNameVoValidName verifica a criação de um novo Name com um valor válido
func (suite *NameTestSuite) TestNewNameVoValidName() {
	name := "ValidName"
	nameVo, err := NewNameVo(name)
	assert.Nil(suite.T(), err, "erro deve ser nil para um nome válido")
	assert.Equal(suite.T(), name, nameVo.Value(), "o valor do nome deve ser igual ao valor fornecido")
}

// TestNewNameVoInvalidName verifica a criação de um novo Name com um valor inválido
func (suite *NameTestSuite) TestNewNameVoInvalidName() {
	name := "Go"
	nameVo, err := NewNameVo(name)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para um nome inválido")
	assert.Equal(suite.T(), ErrInvalidName, err, "erro retornado deve ser ErrInvalidName")
	assert.Nil(suite.T(), nameVo, "Name deve ser nil para um valor de nome inválido")
}

// TestNameValue verifica o método Value da struct Name
func (suite *NameTestSuite) TestNameValue() {
	name := "AnotherName"
	nameVo, _ := NewNameVo(name)
	assert.Equal(suite.T(), name, nameVo.Value(), "o valor do nome deve ser igual ao valor fornecido")
}

// TestNameTestSuite inicia o suite de teste
func TestNameTestSuite(t *testing.T) {
	suite.Run(t, new(NameTestSuite))
}
