package client

import (
	"github.com/imroc/req/v3"
)

func SendRequest(url string) {

	req.EnableForceHTTP1() // Force using HTTP/1.1
	req.MustGet(url)

	// fmt.Println(resp)
}
