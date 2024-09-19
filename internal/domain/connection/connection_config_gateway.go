package connection

import "github.com/google/uuid"

type ConnectionConfigGateway interface {
	FindById(id uuid.UUID) (ConnectionConfigEntity, error)
	FindByUserNameAndHost(userName string, host string) (ConnectionConfigEntity, error)
	Save(entity ConnectionConfigEntity) (ConnectionConfigEntity, error)
	Delete(entity ConnectionConfigEntity) (ConnectionConfigEntity, error)
	ExistsUserNameAndHostAndIdNot(userName string, host string, id uuid.UUID) (bool, error)
}
