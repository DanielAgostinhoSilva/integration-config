package flow

import (
	"github.com/stretchr/testify/mock"
)

// MockPullFlowGateway é a estrutura que implementa o mock da interface PullFlowGateway
type MockPullFlowGateway struct {
	mock.Mock
}

// Save mock do método Save da interface PullFlowGateway
func (m *MockPullFlowGateway) Save(config PullFlowConfig) {
	m.Called(config)
}

// Delete mock do método Delete da interface PullFlowGateway
func (m *MockPullFlowGateway) Delete(config PullFlowConfig) {
	m.Called(config)
}

// FindById mock do método FindById da interface PullFlowGateway
func (m *MockPullFlowGateway) FindById(id string) (PullFlowConfig, error) {
	args := m.Called(id)
	return args.Get(0).(PullFlowConfig), args.Error(1)
}

// ExistOriginPathAndPrefixAndIdNot mock do método ExistOriginPathAndPrefixAndIdNot da interface PullFlowGateway
func (m *MockPullFlowGateway) ExistOriginPathAndPrefixAndIdNot(originPath, prefix, id string) bool {
	args := m.Called(originPath, prefix, id)
	return args.Bool(0)
}

// ExistOriginPathAndSuffixAndIdNot mock do método ExistOriginPathAndSuffixAndIdNot da interface PullFlowGateway
func (m *MockPullFlowGateway) ExistOriginPathAndSuffixAndIdNot(originPath, suffix, id string) bool {
	args := m.Called(originPath, suffix, id)
	return args.Bool(0)
}

// ExistOriginPathAndRegexAndIdNot mock do método ExistOriginPathAndRegexAndIdNot da interface PullFlowGateway
func (m *MockPullFlowGateway) ExistOriginPathAndRegexAndIdNot(originPath, regex, id string) bool {
	args := m.Called(originPath, regex, id)
	return args.Bool(0)
}
