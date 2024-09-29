package connection

type CreateCommand struct {
	Host                    string
	Port                    any
	Username                string
	Password                string
	ConnectionConfigGateway ConnectionConfigGateway
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
