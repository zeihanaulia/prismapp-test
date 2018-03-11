package main

import (
	"flag"
	"log"
	"net/http"

	handlers "prismapp-test/src/websocket/handlers"
	usecase "prismapp-test/src/websocket/usecase"
)

func main() {
	var addr = flag.String("addr", ":7070", "The addr of the  application.")
	flag.Parse() // parse the flags

	r := usecase.NewRoom()
	http.Handle("/realtime", r)
	http.Handle("/", &handlers.TemplateHandler{Filename: "chat.html"})
	go r.Run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
