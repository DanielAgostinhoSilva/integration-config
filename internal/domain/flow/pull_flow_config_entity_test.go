package flow

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// PullFlowTestSuite define a estrutura para o suite de testes da entidade PullFlowConfig
type PullFlowTestSuite struct {
	suite.Suite
	mockGateway *MockPullFlowGateway
}

// SetupTest configura qualquer preparo necessário antes de cada teste
func (suite *PullFlowTestSuite) SetupTest() {
	suite.mockGateway = &MockPullFlowGateway{}
}

// TestNewPullFlowConfigValid testa a criação de um PullFlowConfig com valores válidos
func (suite *PullFlowTestSuite) TestNewPullFlowConfigValid() {
	props := PullFlowProps{
		Id:              uuid.New(),
		ConnectionId:    uuid.New(),
		Name:            "TestName",
		IntegrationType: SFTP,
		OriginPath:      "test/origin/path",
		TargetPath:      "test/target/path",
		PrefixFilter:    "prefix",
		SuffixFilter:    "suffix",
		RegexFilter:     "regex",
	}

	config, err := NewPullFlowConfig(props)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), config)
	assert.Equal(suite.T(), props.Name, config.Name().Value())
	assert.Equal(suite.T(), props.OriginPath, config.OriginPath().Value())
	assert.Equal(suite.T(), props.TargetPath, config.TargetPath().Value())
}

// TestCreatePullFlowConfigValid testa a criação de um PullFlowConfig através do comando CreatePullFlowConfig
func (suite *PullFlowTestSuite) TestCreatePullFlowConfigValid() {
	suite.mockGateway.On("ExistOriginPathAndPrefixAndIdNot", "origin/path", "prefix", mock.Anything).Return(false)
	suite.mockGateway.On("ExistOriginPathAndSuffixAndIdNot", "origin/path", "suffix", mock.Anything).Return(false)
	suite.mockGateway.On("ExistOriginPathAndRegexAndIdNot", "origin/path", "regex", mock.Anything).Return(false)

	command := CreatePullConfigCommand{
		ConnectionId:    uuid.New(),
		Name:            "example name",
		IntegrationType: SFTP,
		OriginPath:      "origin/path",
		TargetPath:      "target/path",
		PrefixFilter:    "prefix",
		SuffixFilter:    "suffix",
		RegexFilter:     "regex",
		PullFlowGateway: suite.mockGateway,
	}

	entity, err := CreatePullFlowConfig(command)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), entity)
	assert.Equal(suite.T(), command.Name, entity.Name().Value())
}

// TestPullFlowConfigUpdateValid testa a atualização de um PullFlowConfig com valores válidos
func (suite *PullFlowTestSuite) TestPullFlowConfigUpdateValid() {
	suite.mockGateway.On("ExistOriginPathAndPrefixAndIdNot", "updated/origin/path", "updatedPrefix", mock.Anything).Return(false)
	suite.mockGateway.On("ExistOriginPathAndSuffixAndIdNot", "updated/origin/path", "updatedSuffix", mock.Anything).Return(false)
	suite.mockGateway.On("ExistOriginPathAndRegexAndIdNot", "updated/origin/path", "updatedRegex", mock.Anything).Return(false)

	props := PullFlowProps{
		Id:              uuid.New(),
		ConnectionId:    uuid.New(),
		Name:            "TestName",
		IntegrationType: SFTP,
		OriginPath:      "test/origin/path",
		TargetPath:      "test/target/path",
		PrefixFilter:    "prefix",
		SuffixFilter:    "suffix",
		RegexFilter:     "regex",
	}
	config, _ := NewPullFlowConfig(props)

	command := UpdatePullConfigCommand{
		ConnectionId:    uuid.New(),
		Name:            "UpdatedName",
		OriginPath:      "updated/origin/path",
		TargetPath:      "updated/target/path",
		PrefixFilter:    "updatedPrefix",
		SuffixFilter:    "updatedSuffix",
		RegexFilter:     "updatedRegex",
		PullFlowGateway: suite.mockGateway,
	}

	err := config.Update(command)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), command.Name, config.Name().Value())
	assert.Equal(suite.T(), command.OriginPath, config.OriginPath().Value())
	assert.Equal(suite.T(), command.TargetPath, config.TargetPath().Value())
	assert.Equal(suite.T(), command.PrefixFilter, config.PrefixFilter())
	assert.Equal(suite.T(), command.SuffixFilter, config.SuffixFilter())
	assert.Equal(suite.T(), command.RegexFilter, config.RegexFilter())
}

// TestValidateFilters testa a validação de filtros em um PullFlowConfig
func (suite *PullFlowTestSuite) TestValidateFilters() {
	props := PullFlowProps{
		Id:              uuid.New(),
		ConnectionId:    uuid.New(),
		Name:            "TestName",
		IntegrationType: SFTP,
		OriginPath:      "test/origin/path",
		TargetPath:      "test/target/path",
		PrefixFilter:    "prefix",
		SuffixFilter:    "suffix",
		RegexFilter:     "regex",
	}
	config, _ := NewPullFlowConfig(props)

	suite.mockGateway.On("ExistOriginPathAndPrefixAndIdNot", "test/origin/path", "prefix", mock.Anything).Return(false)
	suite.mockGateway.On("ExistOriginPathAndSuffixAndIdNot", "test/origin/path", "suffix", mock.Anything).Return(false)
	suite.mockGateway.On("ExistOriginPathAndRegexAndIdNot", "test/origin/path", "regex", mock.Anything).Return(false)

	err := config.ValidateFilters(suite.mockGateway)
	assert.Nil(suite.T(), err)

	suite.mockGateway.On("ExistOriginPathAndPrefixAndIdNot", "test/origin/path", "errorPrefix", mock.Anything).Return(true)
	props.PrefixFilter = "errorPrefix"
	config, _ = NewPullFlowConfig(props)
	err = config.ValidateFilters(suite.mockGateway)

	assert.NotNil(suite.T(), err)
	assert.IsType(suite.T(), &errors.EntityInUseError{}, err)
}

// TestPullFlowTestSuite inicia o suite de testes
func TestPullFlowTestSuite(t *testing.T) {
	suite.Run(t, new(PullFlowTestSuite))
}
