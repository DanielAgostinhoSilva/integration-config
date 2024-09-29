package flow

import (
	"fmt"
	flow "github.com/DanielAgostinhoSilva/integration-config/internal/domain/flow/vo"
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/vo"
	"github.com/google/uuid"
)

type PullFlowProps struct {
	Id              any
	ConnectionId    any
	Name            string
	IntegrationType IntegrationType
	OriginPath      string
	TargetPath      string
	PrefixFilter    string
	SuffixFilter    string
	RegexFilter     string
}

type PullFlowConfig struct {
	id              *vo.ID
	connectionId    *vo.ID
	name            *vo.Name
	integrationType IntegrationType
	originPath      *flow.OriginPath
	targetPath      *flow.TargetPath
	prefixFilter    string
	suffixFilter    string
	regexFilter     string
}

func NewPullFlowConfig(props PullFlowProps) (*PullFlowConfig, error) {
	id, err := vo.NewID(props.Id)
	if err != nil {
		return nil, err
	}
	connectionId, err := vo.NewID(props.ConnectionId)
	if err != nil {
		return nil, err
	}
	name, err := vo.NewNameVo(props.Name)
	if err != nil {
		return nil, err
	}
	originPath, err := flow.NewOriginPathVo(props.OriginPath)
	if err != nil {
		return nil, err
	}
	targetPath, err := flow.NewTargetPathVo(props.TargetPath)
	if err != nil {
		return nil, err
	}

	return &PullFlowConfig{
		id:              id,
		connectionId:    connectionId,
		name:            name,
		integrationType: props.IntegrationType,
		originPath:      originPath,
		targetPath:      targetPath,
		prefixFilter:    props.PrefixFilter,
		suffixFilter:    props.SuffixFilter,
		regexFilter:     props.RegexFilter,
	}, nil
}

func CreatePullFlowConfig(command CreatePullConfigCommand) (*PullFlowConfig, error) {
	props := PullFlowProps{
		Id:              uuid.New(),
		ConnectionId:    command.ConnectionId,
		Name:            command.Name,
		IntegrationType: command.IntegrationType,
		OriginPath:      command.OriginPath,
		TargetPath:      command.TargetPath,
		PrefixFilter:    command.PrefixFilter,
		SuffixFilter:    command.SuffixFilter,
		RegexFilter:     command.RegexFilter,
	}

	entity, err := NewPullFlowConfig(props)
	if err != nil {
		return nil, err
	}

	if err = entity.ValidateFilters(command.PullFlowGateway); err != nil {
		return nil, err
	}

	return entity, nil
}

func (p *PullFlowConfig) Update(command UpdatePullConfigCommand) error {
	connectionId, err := vo.NewID(command.ConnectionId)
	if err != nil {
		return err
	}
	name, err := vo.NewNameVo(command.Name)
	if err != nil {
		return err
	}
	originPath, err := flow.NewOriginPathVo(command.OriginPath)
	if err != nil {
		return err
	}
	targetPath, err := flow.NewTargetPathVo(command.TargetPath)
	if err != nil {
		return err
	}

	p.connectionId = connectionId
	p.name = name
	p.originPath = originPath
	p.targetPath = targetPath
	p.prefixFilter = command.PrefixFilter
	p.suffixFilter = command.SuffixFilter
	p.regexFilter = command.RegexFilter

	return p.ValidateFilters(command.PullFlowGateway)
}

func (p *PullFlowConfig) ValidateFilters(pullFlowGateway PullFlowGateway) error {
	if pullFlowGateway.ExistOriginPathAndPrefixAndIdNot(p.originPath.Value(), p.prefixFilter, p.id.Value().String()) {
		return errors.NewEntityInUseError(fmt.Sprintf("origin path %s and prefix %s already in use", p.originPath.Value(), p.prefixFilter))
	}

	if pullFlowGateway.ExistOriginPathAndSuffixAndIdNot(p.originPath.Value(), p.suffixFilter, p.id.Value().String()) {
		return errors.NewEntityInUseError(fmt.Sprintf("origin path %s and suffix %s already in use", p.originPath.Value(), p.suffixFilter))
	}

	if pullFlowGateway.ExistOriginPathAndRegexAndIdNot(p.originPath.Value(), p.regexFilter, p.id.Value().String()) {
		return errors.NewEntityInUseError(fmt.Sprintf("origin path %s and regex %s already in use", p.originPath.Value(), p.regexFilter))
	}
	return nil
}

func (p *PullFlowConfig) Id() *vo.ID {
	return p.id
}

func (p *PullFlowConfig) ConnectionId() *vo.ID {
	return p.connectionId
}

func (p *PullFlowConfig) Name() *vo.Name {
	return p.name
}

func (p *PullFlowConfig) IntegrationType() IntegrationType {
	return p.integrationType
}

func (p *PullFlowConfig) OriginPath() *flow.OriginPath {
	return p.originPath
}

func (p *PullFlowConfig) TargetPath() *flow.TargetPath {
	return p.targetPath
}

func (p *PullFlowConfig) PrefixFilter() string {
	return p.prefixFilter
}

func (p *PullFlowConfig) SuffixFilter() string {
	return p.suffixFilter
}

func (p *PullFlowConfig) RegexFilter() string {
	return p.regexFilter
}
