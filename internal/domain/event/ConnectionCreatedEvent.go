package event

import (
	"github.com/DanielAgostinhoSilva/go-domain-event/events"
	"github.com/DanielAgostinhoSilva/integration-config/internal/domain/connection"
	"github.com/google/uuid"
	"time"
)

type ConnectionCreatedEventData struct {
	Id       string
	host     string
	port     int
	userName string
	password string
}

type ConnectionCreatedEvent struct {
	events.BaseEvent
}

func NewConnectionCreatedEvent(entity *connection.ConnectionConfigEntity) *ConnectionCreatedEvent {
	data := ConnectionCreatedEventData{
		Id:       entity.Id().Value().String(),
		host:     entity.Host().Value(),
		port:     entity.Port().Value(),
		userName: entity.UserName().Value(),
		password: entity.Password().Value(),
	}

	return &ConnectionCreatedEvent{
		BaseEvent: events.BaseEvent{
			ID:            uuid.New().String(),
			Type:          "ConnectionCreatedEvent",
			AggregateID:   data.Id,
			AggregateType: "ConnectionConfig",
			Timestamp:     time.Now(),
			Version:       1,
			Data:          data,
		},
	}
}
