package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func UDPClientDial(host string, port int) {
	addr := fmt.Sprint(host, ":", port)

	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	n, err := conn.Write([]byte("Ping"))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// if len(buf) != n {
	// 	log.Printf("data size is %d, but sent data size is %d", len(buf), n)
	// }

	recvBuf := make([]byte, 1024)

	n, err = conn.Read(recvBuf)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	log.Printf("Received data: %s", string(recvBuf[:n]))
}

func main() {
	UDPClientDial("127.0.0.1", 20001)
}
