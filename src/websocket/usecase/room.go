package usecase

import (
	"log"
	"net/http"

	gateway "prismapp-test/src/websocket/gateways"

	"github.com/gorilla/websocket"
)

type room struct {
	// forward untuk menyimpan pesan didalam channel yang nantinya akan diteruskan ke client
	forward chan []byte
	// join sebagai channel jika ingin bergabung ke dalam room chat
	join chan *client
	// leave channel untuk client yang ingin meninggalkan room
	leave chan *client
	// clients dalam map / menahan seluruh client yang ada didalam ruangan
	clients map[*client]bool

	//gateway to database
	roomGateway gateway.Room
}

// NewRoom initialize access object
func NewRoom(roomGateway gateway.Room) *room {
	return &room{
		forward:     make(chan []byte),
		join:        make(chan *client),
		leave:       make(chan *client),
		clients:     make(map[*client]bool),
		roomGateway: roomGateway,
	}
}

func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
