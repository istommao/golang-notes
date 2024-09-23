package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

type UDPClient struct {
	conn net.Conn
}

func NewUDPClient(host string, port int) *UDPClient {
	addr := fmt.Sprint(host, ":", port)

	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	client := &UDPClient{conn: conn}
	return client
}

func (client *UDPClient) Close() error {
	return client.conn.Close()
}

func (client *UDPClient) Write(data string) {
	n, err := client.conn.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// if len(buf) != n {
	//     log.Printf("data size is %d, but sent data size is %d", len(buf), n)
	// }

	recvBuf := make([]byte, 1024)

	n, err = client.conn.Read(recvBuf)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	log.Printf("Received data: %s", string(recvBuf[:n]))
}
