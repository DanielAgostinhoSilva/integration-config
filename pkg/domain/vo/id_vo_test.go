package vo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

// IDTestSuite é a estrutura da suite de testes para o tipo ID.
type IDTestSuite struct {
	suite.Suite
}

// SetupSuite é executado uma vez antes de todos os testes na suite.
func (suite *IDTestSuite) SetupSuite() {
	// Qualquer configuração de inicialização, se necessário
}

// TestNewIDFromString testa a criação de ID a partir de uma string válida.
func (suite *IDTestSuite) TestNewIDFromString() {
	idString := "550e8400-e29b-41d4-a716-446655440000"
	id, err := NewID(idString)
	suite.Nil(err, "Erro deve ser nulo")
	suite.NotNil(id, "O ID não deve ser nulo")
	suite.Equal(idString, id.Value().String(), "Os UUIDs devem ser iguais")
}

// TestNewIDFromInvalidString testa a criação de ID a partir de uma string inválida.
func (suite *IDTestSuite) TestNewIDFromInvalidString() {
	idString := "invalid-uuid-string"
	id, err := NewID(idString)
	suite.NotNil(err, "Erro deve estar presente")
	suite.Nil(id, "O ID deve ser nulo")
}

// TestNewIDFromUUID testa a criação de ID a partir de um valor uuid.UUID.
func (suite *IDTestSuite) TestNewIDFromUUID() {
	uuidValue := uuid.New()
	id, err := NewID(uuidValue)
	suite.Nil(err, "Erro deve ser nulo")
	suite.NotNil(id, "O ID não deve ser nulo")
	suite.Equal(uuidValue, id.Value(), "Os UUIDs devem ser iguais")
}

// TestNewIDFromUnsupportedType testa a criação de ID a partir de um tipo não suportado.
func (suite *IDTestSuite) TestNewIDFromUnsupportedType() {
	unsupportedValue := 12345
	id, err := NewID(unsupportedValue)
	suite.NotNil(err, "Erro deve estar presente")
	suite.Nil(id, "O ID deve ser nulo")
}

// TestIDTestSuite executa a suite de testes.
func TestIDTestSuite(t *testing.T) {
	suite.Run(t, new(IDTestSuite))
}
