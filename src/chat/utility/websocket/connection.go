package websocket

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

// Connection to websocket
func Connection(addr, path string) *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: addr, Path: path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	//defer c.Close()
	return c
}
