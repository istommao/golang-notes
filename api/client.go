package api

type BaseClient interface {
	// init config
	Init(conf ...interface{}) error

	Send(data []byte) error

	Stop() error
}

type HTTPClient interface {
	BaseClient
}

type TCPClient interface {
	BaseClient
}

type UDPClient interface {
	BaseClient
}

type QUICClient interface {
	BaseClient
}

type WebsocketClient interface {
	BaseClient
}
