package main

import (
	"log"
	"net"
	"os"
)

func CreateUDPServer(host string, port int) *net.UDPConn {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}
	updLn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	log.Println("Starting udp server...")

	go func() {
		for {
			n, addr, err := updLn.ReadFromUDP(buf)
			if err != nil {
				log.Fatalln(err)
				os.Exit(1)
			}

			go func() {
				log.Printf("Reciving data: %s from %s", string(buf[:n]), addr.String())

				log.Printf("Sending data..")
				updLn.WriteTo([]byte("Pong"), addr)
				log.Printf("Complete Sending data..")
			}()
		}
	}()

	return updLn
}
