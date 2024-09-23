package client

import (
	"testing"
)

func TestHTTPClient(t *testing.T) {
	SendRequest("https://httpbin.org/")
}
