package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUDPClient(t *testing.T) {
	client := new(DtUDPClient)
	err := client.Init("127.0.0.1", 20001)
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	client.Send([]byte("hello"))

	assert.Equal(t, client.host, "127.0.0.1", "udp host check.")
	fmt.Println(client.port)
	assert.Equal(t, client.port, 20001, "udp port check.")
}
