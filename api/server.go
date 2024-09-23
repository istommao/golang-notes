package api

type BaseServer interface {
	// init config
	Init(conf ...interface{}) error

	Start() error

	Stop() error
}

type HTTPServer interface {
	BaseServer
}

type WebsocketServer interface {
	BaseServer
}

type TCPServer interface {
	BaseServer

	ReadLoop() error
}

type UDPServer interface {
	BaseServer

	// ReadLoop
	ReadLoop() error
}

type QUICServer interface {
	BaseServer
}
