package main

// import "time"

func main() {
	client := NewUDPClient("127.0.0.1", 30000)

	// client.Write("ping")
	// client.Write("ping")
	// client.Write("ping")
	client.Write("ping")
	// time.Sleep(1 * time.Second)
}
