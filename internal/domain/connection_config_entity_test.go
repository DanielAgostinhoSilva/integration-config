package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConnectionConfigEntityTestSuite struct {
	suite.Suite
}

// TestNewConnectionConfigEntity verifica a criação de uma nova ConnectionConfigEntity com dados válidos
func (suite *ConnectionConfigEntityTestSuite) TestNewConnectionConfigEntity() {
	id := "550e8400-e29b-41d4-a716-446655440000"
	host := "192.168.1.10"
	port := 8080
	password := "Valid1@Password"

	entity, err := NewConnectionConfigEntity(id, host, port, password)
	assert.Nil(suite.T(), err, "erro deve ser nil para dados válidos")

	assert.Equal(suite.T(), id, entity.Id().Value().String(), "ID deve ser igual ao valor fornecido")
	assert.Equal(suite.T(), host, entity.Host().Value(), "Host deve ser igual ao valor fornecido")
	assert.Equal(suite.T(), port, entity.Port().Value(), "Port deve ser igual ao valor fornecido")
	assert.Equal(suite.T(), password, entity.Password().Value(), "Password deve ser igual ao valor fornecido")
}

// TestNewConnectionConfigEntityInvalidID verifica a criação de ConnectionConfigEntity com ID inválido
func (suite *ConnectionConfigEntityTestSuite) TestNewConnectionConfigEntityInvalidID() {
	host := "localhost"
	port := 8080
	password := "Valid1@Password"

	entity, err := NewConnectionConfigEntity(nil, host, port, password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para ID inválido")
	assert.Nil(suite.T(), entity, "entity deve ser nil para ID inválido")
}

// TestNewConnectionConfigEntityInvalidHost verifica a criação de ConnectionConfigEntity com host inválido
func (suite *ConnectionConfigEntityTestSuite) TestNewConnectionConfigEntityInvalidHost() {
	id := "valid-id"
	port := 8080
	password := "Valid1@Password"

	entity, err := NewConnectionConfigEntity(id, "invalid-host", port, password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para host inválido")
	assert.Nil(suite.T(), entity, "entity deve ser nil para host inválido")
}

// TestSetHost verifica a atualização do host com valor válido
func (suite *ConnectionConfigEntityTestSuite) TestSetHost() {
	id := "550e8400-e29b-41d4-a716-446655440000"
	host := "192.168.1.10"
	port := 8080
	password := "Valid1@Password"

	entity, _ := NewConnectionConfigEntity(id, host, port, password)

	newHost := "127.0.0.1"
	err := entity.SetHost(newHost)
	assert.Nil(suite.T(), err, "erro deve ser nil ao atualizar o host com valor válido")
	assert.Equal(suite.T(), newHost, entity.Host().Value(), "Host deve ser igual ao novo valor fornecido")
}

// TestSetPort verifica a atualização da porta com valor válido
func (suite *ConnectionConfigEntityTestSuite) TestSetPort() {
	id := "550e8400-e29b-41d4-a716-446655440000"
	host := "192.168.1.10"
	port := 8080
	password := "Valid1@Password"

	entity, _ := NewConnectionConfigEntity(id, host, port, password)

	newPort := 9090
	err := entity.SetPort(newPort)
	assert.Nil(suite.T(), err, "erro deve ser nil ao atualizar a porta com valor válido")
	assert.Equal(suite.T(), newPort, entity.Port().Value(), "Port deve ser igual ao novo valor fornecido")
}

// TestSetPassword verifica a atualização da senha com valor válido
func (suite *ConnectionConfigEntityTestSuite) TestSetPassword() {
	id := "550e8400-e29b-41d4-a716-446655440000"
	host := "192.168.1.10"
	port := 8080
	password := "Valid1@Password"

	entity, _ := NewConnectionConfigEntity(id, host, port, password)

	newPassword := "NewValid1@Password"
	err := entity.SetPassword(newPassword)
	assert.Nil(suite.T(), err, "erro deve ser nil ao atualizar a senha com valor válido")
	assert.Equal(suite.T(), newPassword, entity.Password().Value(), "Password deve ser igual ao novo valor fornecido")
}

// TestConnectionConfigEntityTestSuite inicia o suite de teste
func TestConnectionConfigEntityTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionConfigEntityTestSuite))
}
