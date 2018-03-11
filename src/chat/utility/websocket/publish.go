package websocket

import (
	"log"
	gateway "prismapp-test/src/chat/gateway"

	"github.com/gorilla/websocket"
)

type publish struct {
	ws *websocket.Conn
}

// NewPublishGateway initial to access this object
func NewPublishGateway(conn *websocket.Conn) gateway.Websocket {
	return &publish{
		ws: conn,
	}
}

func (p *publish) Publish(message string) {
	go func(message string) {

		err := p.ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("write:", err)
			return
		}

	}(message)
}
