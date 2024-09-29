package connection

import (
	"context"
	"github.com/DanielAgostinhoSilva/go-domain-event/events"
	"github.com/DanielAgostinhoSilva/integration-config/internal/domain/connection"
	"github.com/DanielAgostinhoSilva/integration-config/internal/domain/event"
)

type CreateConnectionInput struct {
	Host     string
	Port     any
	Username string
	Password string
}

type CreateConnectionOutput struct {
	Id       string
	Host     string
	Port     any
	Username string
}

type CreateConnectionUseCase struct {
	ConnectionGateway connection.ConnectionConfigGateway
	EventDispatcher   events.EventDispatcher
}

func NewCreateConnectionUseCase(connectionGateway connection.ConnectionConfigGateway, eventDispatcher events.EventDispatcher) *CreateConnectionUseCase {
	return &CreateConnectionUseCase{ConnectionGateway: connectionGateway, EventDispatcher: eventDispatcher}
}

func (c *CreateConnectionUseCase) Execute(ctx context.Context, input CreateConnectionInput) (*CreateConnectionOutput, error) {
	entity, err := connection.CreateConnectionConfigEntity(connection.CreateCommand{
		ConnectionConfigGateway: c.ConnectionGateway,
		Host:                    input.Host,
		Port:                    input.Port,
		Username:                input.Username,
		Password:                input.Password,
	})
	if err != nil {
		return nil, err
	}
	entity, err = c.ConnectionGateway.Save(entity)
	if err != nil {
		return nil, err
	}

	c.EventDispatcher.Dispatch(ctx, event.NewConnectionCreatedEvent(entity))

	return &CreateConnectionOutput{
		Id:       entity.Id().Value().String(),
		Host:     entity.Host().Value(),
		Port:     entity.Port().Value(),
		Username: entity.UserName().Value(),
	}, nil
}
