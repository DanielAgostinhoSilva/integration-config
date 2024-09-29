package connection

import (
	"context"
	"fmt"
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/vo"
	"github.com/google/uuid"
)

type ConnectionConfigEntity struct {
	id       *vo.ID
	host     *vo.NetworkAddress
	port     *vo.NetworkPort
	userName *vo.Name
	password *vo.PasswordVo
}

func NewConnectionConfigEntity(
	id any,
	host string,
	port any,
	name string,
	password string,
) (*ConnectionConfigEntity, error) {
	idVo, err := vo.NewID(id)
	if err != nil {
		return nil, err
	}
	hostVo, err := vo.NewNetworkAddress(context.Background(), host)
	if err != nil {
		return nil, err
	}
	portVo, err := vo.NewNetworkPortVo(port)
	if err != nil {
		return nil, err
	}
	userNameVo, err := vo.NewNameVo(name)
	if err != nil {
		return nil, err
	}
	passwordVo, err := vo.NewPasswordVo(password)
	if err != nil {
		return nil, err
	}
	return &ConnectionConfigEntity{
		id:       idVo,
		host:     hostVo,
		port:     portVo,
		userName: userNameVo,
		password: passwordVo,
	}, nil
}

func CreateConnectionConfigEntity(command CreateCommand) (*ConnectionConfigEntity, error) {
	entity, err := NewConnectionConfigEntity(
		uuid.New(),
		command.Host,
		command.Port,
		command.Username,
		command.Password,
	)
	if err != nil {
		return nil, err
	}
	exist, err := command.ConnectionConfigGateway.ExistsUserNameAndHostAndIdNot(
		entity.userName.Value(),
		entity.host.Value(),
		entity.id.Value(),
	)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.NewEntityInUseError(fmt.Sprintf("user name %s and host %s already exists", entity.userName.Value(), entity.host.Value()))
	}

	return entity, nil
}

func (c *ConnectionConfigEntity) UpdateHost(command UpdateHostCommand) error {
	host, err := vo.NewNetworkAddress(context.Background(), command.Host)
	if err != nil {
		return err
	}

	exist, err := command.ConnectionConfigEntity.ExistsUserNameAndHostAndIdNot(
		c.userName.Value(),
		host.Value(),
		c.id.Value(),
	)
	if err != nil {
		return err
	}
	if exist {
		return errors.NewEntityInUseError(fmt.Sprintf("user name %s and host %s already exists", c.userName.Value(), command.Host))
	}
	c.host = host
	return nil
}

func (c *ConnectionConfigEntity) UpdatePort(command UpdatePortCommand) error {
	port, err := vo.NewNetworkPortVo(command.Port)
	if err != nil {
		return err
	}
	c.port = port
	return nil
}

func (c *ConnectionConfigEntity) UpdateUserName(command UpdateUsernameCommand) error {
	user, err := vo.NewNameVo(command.Username)
	if err != nil {
		return err
	}
	exist, err := command.ConnectionConfigEntity.ExistsUserNameAndHostAndIdNot(
		user.Value(),
		c.host.Value(),
		c.id.Value(),
	)
	if err != nil {
		return err
	}
	if exist {
		return errors.NewEntityInUseError(fmt.Sprintf("user name %s and host %s already exists", user.Value(), c.host.Value()))
	}
	c.userName = user
	return nil
}

func (c *ConnectionConfigEntity) UpdatePassword(command UpdatePasswordCommand) error {
	password, err := vo.NewPasswordVo(command.Password)
	if err != nil {
		return err
	}
	c.password = password
	return nil
}

func (c *ConnectionConfigEntity) Id() *vo.ID {
	return c.id
}

func (c *ConnectionConfigEntity) Host() *vo.NetworkAddress {
	return c.host
}

func (c *ConnectionConfigEntity) Port() *vo.NetworkPort {
	return c.port
}

func (c *ConnectionConfigEntity) UserName() *vo.Name {
	return c.userName
}
func (c *ConnectionConfigEntity) Password() *vo.PasswordVo {
	return c.password
}
