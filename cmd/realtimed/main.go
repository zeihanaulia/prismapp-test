package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "prismapp-test/src/websocket/handlers"
	usecase "prismapp-test/src/websocket/usecase"

	database "prismapp-test/src/websocket/utility/mysql"

	"github.com/joho/godotenv"
)

func main() {
	var addr = flag.String("addr", ":7070", "The addr of the  application.")
	flag.Parse() // parse the flags

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	connection := database.Connection(os.Getenv("CONNECTION_STRING"))
	roomGateway := database.NewRoomGateway(connection)

	r := usecase.NewRoom(roomGateway)
	http.Handle("/realtime", r)
	http.Handle("/", &handlers.TemplateHandler{Filename: "chat.html"})
	go r.Run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
