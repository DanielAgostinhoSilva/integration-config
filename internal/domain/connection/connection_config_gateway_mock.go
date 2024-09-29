package connection

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockConnectionConfigGateway é o mock da interface ConnectionConfigGateway
type MockConnectionConfigGateway struct {
	mock.Mock
}

func (m *MockConnectionConfigGateway) FindById(id uuid.UUID) (ConnectionConfigEntity, error) {
	args := m.Called(id)
	return args.Get(0).(ConnectionConfigEntity), args.Error(1)
}

func (m *MockConnectionConfigGateway) FindByUserNameAndHost(userName string, host string) (ConnectionConfigEntity, error) {
	args := m.Called(userName, host)
	return args.Get(0).(ConnectionConfigEntity), args.Error(1)
}

func (m *MockConnectionConfigGateway) Save(entity *ConnectionConfigEntity) (*ConnectionConfigEntity, error) {
	args := m.Called(entity)
	return args.Get(0).(*ConnectionConfigEntity), args.Error(1)
}

func (m *MockConnectionConfigGateway) Delete(entity *ConnectionConfigEntity) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockConnectionConfigGateway) ExistsUserNameAndHostAndIdNot(userName string, host string, id uuid.UUID) (bool, error) {
	args := m.Called(userName, host, id)
	return args.Bool(0), args.Error(1)
}
