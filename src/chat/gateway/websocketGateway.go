package gateway

type Websocket interface {
	Publish(message string)
}
