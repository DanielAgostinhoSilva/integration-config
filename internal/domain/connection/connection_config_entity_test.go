package connection

import (
	domainErrors "github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ConnectionConfigEntityTestSuite estrutura para os testes de ConnectionConfigEntity
type ConnectionConfigEntityTestSuite struct {
	suite.Suite
	mockEntity *MockConnectionConfigGateway
}

// SetupTest configura antes de cada teste
func (suite *ConnectionConfigEntityTestSuite) SetupTest() {
	suite.mockEntity = &MockConnectionConfigGateway{}
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

// TestUpdateHost verifies the UpdateHost method of ConnectionConfigEntity updates the host correctly.
func (suite *ConnectionConfigEntityTestSuite) TestUpdateHost() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "192.168.1.1", mock.Anything).Return(false, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdateHost(UpdateHostCommand{Host: "192.168.1.1", ConnectionConfigEntity: suite.mockEntity})

	assert.Nil(suite.T(), err, "err deve ser nulo")
	assert.Equal(suite.T(), "192.168.1.1", entity.Host().Value(), "host deve ser atualizado")
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestUpdateHostEntityExists verifies that the UpdateHost method returns an error when the new host already exists for a given username.
func (suite *ConnectionConfigEntityTestSuite) TestUpdateHostEntityExists() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "192.168.1.1", mock.Anything).Return(true, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdateHost(UpdateHostCommand{Host: "192.168.1.1", ConnectionConfigEntity: suite.mockEntity})

	assert.NotNil(suite.T(), err, "erro deve ser retornado quando o usuário e host já existem")
	assert.Equal(suite.T(), domainErrors.NewEntityInUseError("user name ValidName and host 192.168.1.1 already exists"), err)
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestUpdateHostInvalidHost verifies that the UpdateHost method returns an error when given an invalid host.
func (suite *ConnectionConfigEntityTestSuite) TestUpdateHostInvalidHost() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdateHost(UpdateHostCommand{Host: "%ˆ&**", ConnectionConfigEntity: suite.mockEntity})

	assert.NotNil(suite.T(), err, "erro deve ser retornado quando o host estiver invalido")
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestUpdatePort verifies that the UpdatePort method of ConnectionConfigEntity updates the port correctly.
func (suite *ConnectionConfigEntityTestSuite) TestUpdatePort() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)

	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdatePort(UpdatePortCommand{Port: 9090})

	assert.Nil(suite.T(), err, "err should be nil")
	assert.Equal(suite.T(), 9090, entity.Port().Value(), "port should be updated")
}

func (suite *ConnectionConfigEntityTestSuite) TestUpdatePortInvalidPort() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)

	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdatePort(UpdatePortCommand{Port: "invalid-port"})

	assert.NotNil(suite.T(), err, "error should be returned when the port is invalid")
}

// TestUpdatePassword verifies the UpdatePassword method of ConnectionConfigEntity updates the password correctly.
func (suite *ConnectionConfigEntityTestSuite) TestUpdatePassword() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)

	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdatePassword(UpdatePasswordCommand{Password: "Valid2@Password"})

	assert.Nil(suite.T(), err, "err should be nil")
	assert.Equal(suite.T(), "Valid2@Password", entity.Password().Value(), "password should be updated")
}

// TestUpdatePasswordInvalidPassword verifies the UpdatePassword method returns an error when given an invalid password.
func (suite *ConnectionConfigEntityTestSuite) TestUpdatePasswordInvalidPassword() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdatePassword(UpdatePasswordCommand{Password: "invalid-password"})

	assert.NotNil(suite.T(), err, "error should be returned when the password is invalid")
}

// TestUpdateUserName verifies the UpdateUserName method of ConnectionConfigEntity updates the username correctly.
func (suite *ConnectionConfigEntityTestSuite) TestUpdateUserName() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "NewName", "localhost", mock.Anything).Return(false, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdateUserName(UpdateUsernameCommand{Username: "NewName", ConnectionConfigEntity: suite.mockEntity})

	assert.Nil(suite.T(), err, "err should be nil")
	assert.Equal(suite.T(), "NewName", entity.UserName().Value(), "username should be updated")
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestUpdateUserNameEntityExists verifies that the UpdateUserName method returns an error when the new username already exists for a given host.
func (suite *ConnectionConfigEntityTestSuite) TestUpdateUserNameEntityExists() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ExistingName", "localhost", mock.Anything).Return(true, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdateUserName(UpdateUsernameCommand{Username: "ExistingName", ConnectionConfigEntity: suite.mockEntity})

	assert.NotNil(suite.T(), err, "error should be returned when the username already exists")
	assert.Equal(suite.T(), domainErrors.NewEntityInUseError("user name ExistingName and host localhost already exists"), err)
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestUpdateUserNameInvalidName verifies the UpdateUserName method returns an error when given an invalid name.
func (suite *ConnectionConfigEntityTestSuite) TestUpdateUserNameInvalidName() {
	suite.mockEntity.On("ExistsUserNameAndHostAndIdNot", "ValidName", "localhost", mock.Anything).Return(false, nil)
	command := CreateCommand{
		Host:                   "localhost",
		Port:                   8080,
		Username:               "ValidName",
		Password:               "Valid1@Password",
		ConnectionConfigEntity: suite.mockEntity,
	}

	entity, _ := CreateConnectionConfigEntity(command)

	err := entity.UpdateUserName(UpdateUsernameCommand{Username: "ab", ConnectionConfigEntity: suite.mockEntity})

	assert.NotNil(suite.T(), err, "error should be returned when the name is invalid")
	suite.mockEntity.AssertExpectations(suite.T())
}

// TestConnectionConfigEntityTestSuite inicia o suite de teste
func TestConnectionConfigEntityTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionConfigEntityTestSuite))
}
