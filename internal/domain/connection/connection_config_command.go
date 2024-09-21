package connection

type CreateCommand struct {
	Host                   string
	Port                   any
	Username               string
	Password               string
	ConnectionConfigEntity ConnectionConfigGateway
}

type UpdateHostCommand struct {
	Host                   string
	ConnectionConfigEntity ConnectionConfigGateway
}

type UpdatePortCommand struct {
	Port any
}

type UpdateUsernameCommand struct {
	Username               string
	ConnectionConfigEntity ConnectionConfigGateway
}

type UpdatePasswordCommand struct {
	Password string
}
