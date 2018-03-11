package usecase

import (
	"log"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {

	// socket akan menahan koneksi ke websocket agar bisa berkomunikasi dengan klien
	socket *websocket.Conn

	// send adalah channle dimana pesan akan dikirim dan membuat antrian yang
	// nantinnya akan dikirim ke klien melalui socket
	send chan []byte

	// room adalah tempat dimana para pengguna bisa saling berkomunikasi
	// semua komunikasi akan ada didalam room
	room *room
}

// read adalah method yang dignakan untuk membaca didalam socket dengan menggunakan method ReadMessage
// dan nantinya akan meneruskan pesan ke channel forward yang ada di room sehingga pengguna akan mendapat
// pesan yang dikirimkan pengguna lainnya
func (c *client) read() {
	defer c.socket.Close() // pencegahan race condition
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		log.Println("recieve: ", string(msg))
		c.room.forward <- msg
	}
}

// write adalah method yang akan digunakan pengguna untuk menulis pesan
// yang ada diladalam chanel send secara terus menerus
func (c *client) write() {
	defer c.socket.Close() // pencegahan race condition
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

func (c *client) Publish(message string) {
	defer c.socket.Close()
	for {
		if message != "" {
			err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				return
			}
		}
	}
}
