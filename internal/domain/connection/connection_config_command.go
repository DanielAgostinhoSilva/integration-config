package connection

type CreateCommand struct {
	Host                   string
	Port                   any
	Username               string
	Password               string
	ConnectionConfigEntity ConnectionConfigGateway
}
