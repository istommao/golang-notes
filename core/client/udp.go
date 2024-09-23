package client

import (
	"errors"
	"fmt"
	"net"
	"tmax/core"
)

type DtUDPClient struct {
	host string
	port int

	conn net.Conn
}

func NewUDPClient() DtUDPClient {
	return DtUDPClient{}
}

func (c *DtUDPClient) GetUID() string {
	addr := c.conn.LocalAddr().String()

	return core.GetMD5Hash(addr)
}

func (c *DtUDPClient) Init(conf ...interface{}) error {
	var (
		host string
		port int
		ok   bool
	)
	if len(conf) != 2 {
		return errors.New("invalid config params")
	}

	if host, ok = conf[0].(string); !ok {
		return errors.New("invalid config param host")
	}

	if port, ok = conf[1].(int); !ok {
		return errors.New("invalid config param port")
	}

	c.host = host
	c.port = port

	addr := fmt.Sprint(host, ":", port)
	_, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("ResolveUDPAddr ===>", err.Error())

		return err
	}

	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Println("===>", err.Error())
		return err
	}

	c.conn = conn

	return nil
}

func (c *DtUDPClient) Send(data []byte) error {
	_, err := c.conn.Write(data)
	return err
}

func (c *DtUDPClient) Stop() error {
	return c.conn.Close()
}
