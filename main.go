// package main

// import "github.com/gofiber/fiber/v2"

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		return ctx.SendString("Hello, World 👋!")
// 	})

// 	app.Listen(":4000")
// }

package main

import (
	"fmt"
	api "tmax/api"
	core "tmax/core/client"
)

type ClientMap struct {
	clients map[string]api.UDPClient
}

func (c *ClientMap) Get(uid string) api.UDPClient {
	if client, ok := c.clients[uid]; ok {
		return client
	} else {
		return nil
	}
}

func main() {
	client := new(core.DtUDPClient)
	err := client.Init("127.0.0.1", 20001)
	if err != nil {
		panic(err)
	}

	xmap := new(ClientMap)
	xmap.clients = map[string]api.UDPClient{client.GetUID(): client}

	k := xmap.Get(client.GetUID())

	fmt.Println(">>>", xmap)

	err = k.Send([]byte("kkzh中文🀄️k"))

	if err != nil {
		fmt.Println(">>>", err.Error())
	}
}
