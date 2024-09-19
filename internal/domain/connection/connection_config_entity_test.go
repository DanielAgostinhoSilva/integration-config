package connection

import (
	"testing"

	domainErrors "github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// ConnectionConfigEntityTestSuite estrutura para os testes de ConnectionConfigEntity
type ConnectionConfigEntityTestSuite struct {
	suite.Suite
	mockEntity *MockConnectionConfigGateway
}

// TestNewConnectionConfigEntity verifica a criação de uma nova ConnectionConfigEntity com dados válidos
func (suite *ConnectionConfigEntityTestSuite) TestNewConnectionConfigEntity() {
	id := uuid.New()
	host := "192.168.1.1"
	port := 8080
	username := "ValidName"
	password := "Valid1@Password"

	entity, err := NewConnectionConfigEntity(id, host, port, username, password)
	assert.Nil(suite.T(), err, "erro deve ser nil para dados válidos")

	assert.Equal(suite.T(), id.String(), entity.Id().Value().String(), "ID deve ser igual ao valor fornecido")
	assert.Equal(suite.T(), host, entity.Host().Value(), "Host deve ser igual ao valor fornecido")
	assert.Equal(suite.T(), port, entity.Port().Value(), "Port deve ser igual ao valor fornecido")
	assert.Equal(suite.T(), username, entity.userName.Value(), "Username deve ser igual ao valor fornecido")
	assert.Equal(suite.T(), password, entity.Password().Value(), "Password deve ser igual ao valor fornecido")
}

// TestNewConnectionConfigEntityInvalidID verifica a criação de ConnectionConfigEntity com ID inválido
func (suite *ConnectionConfigEntityTestSuite) TestNewConnectionConfigEntityInvalidID() {
	host := "localhost"
	port := 8080
	username := "ValidName"
	password := "Valid1@Password"

	entity, err := NewConnectionConfigEntity(nil, host, port, username, password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para ID inválido")
	assert.Nil(suite.T(), entity, "entity deve ser nil para ID inválido")
}

// TestNewConnectionConfigEntityInvalidHost verifica a criação de ConnectionConfigEntity com host inválido
func (suite *ConnectionConfigEntityTestSuite) TestNewConnectionConfigEntityInvalidHost() {
	id := uuid.New()
	port := 8080
	username := "ValidName"
	password := "Valid1@Password"

	entity, err := NewConnectionConfigEntity(id, "invalid-host", port, username, password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para host inválido")
	assert.Nil(suite.T(), entity, "entity deve ser nil para host inválido")
}

// TestCreateConnectionConfigEntity verifica a criação de ConnectionConfigEntity usando CreateCommand com dados válidos
func (suite *ConnectionConfigEntityTestSuite) TestCreateConnectionConfigEntity() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, err := CreateConnectionConfigEntity(command)
	assert.Nil(suite.T(), err, "erro deve ser nil para dados válidos")
	assert.NotNil(suite.T(), entity, "entity não deve ser nil para dados válidos")
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestCreateConnectionConfigEntityExists verifica a criação de ConnectionConfigEntity usando CreateCommand quando usuário e host já existem
func (suite *ConnectionConfigEntityTestSuite) TestCreateConnectionConfigEntityExists() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(true, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, err := CreateConnectionConfigEntity(command)
	assert.NotNil(suite.T(), err, "erro deve ser retornado quando o usuário e host já existem")
	assert.Nil(suite.T(), entity, "entity deve ser nil quando o usuário e host já existem")
	assert.Equal(suite.T(), domainErrors.NewEntityInUseError("user name ValidName and host localhost already exists"), err)
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestConnectionConfigEntityTestSuite inicia o suite de teste
func TestConnectionConfigEntityTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionConfigEntityTestSuite))
}

// SetupTest configura antes de cada teste
func (suite *ConnectionConfigEntityTestSuite) SetupTest() {
	suite.mockEntity = &MockConnectionConfigGateway{}
}
