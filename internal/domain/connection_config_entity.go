package domain

import (
	"context"
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/vo"
)

type ConnectionConfigEntity struct {
	id       *vo.ID
	host     *vo.NetworkAddress
	port     *vo.NetworkPort
	password *vo.PasswordVo
}

func NewConnectionConfigEntity(
	id any,
	host string,
	port any,
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
	passwordVo, err := vo.NewPasswordVo(password)
	if err != nil {
		return nil, err
	}
	return &ConnectionConfigEntity{id: idVo, host: hostVo, port: portVo, password: passwordVo}, nil
}

func (c *ConnectionConfigEntity) Id() *vo.ID {
	return c.id
}

func (c *ConnectionConfigEntity) Host() *vo.NetworkAddress {
	return c.host
}

func (c *ConnectionConfigEntity) SetHost(host string) error {
	if addr, err := vo.NewNetworkAddress(context.Background(), host); err != nil {
		return err
	} else {
		c.host = addr
		return nil
	}
}

func (c *ConnectionConfigEntity) Port() *vo.NetworkPort {
	return c.port
}

func (c *ConnectionConfigEntity) SetPort(port any) error {
	if portVo, err := vo.NewNetworkPortVo(port); err != nil {
		return err
	} else {
		c.port = portVo
		return nil
	}
}

func (c *ConnectionConfigEntity) Password() *vo.PasswordVo {
	return c.password
}

func (c *ConnectionConfigEntity) SetPassword(password string) error {
	if pass, err := vo.NewPasswordVo(password); err != nil {
		return err
	} else {
		c.password = pass
		return nil
	}
}
