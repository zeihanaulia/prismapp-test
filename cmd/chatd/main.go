package main

import (
	"fmt"
	"os"
	server "prismapp-test/src/chat/server"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	var app = func(srv *server.Server) {
		srv.Address = os.Getenv("ADDRESS")
		srv.WsHost = os.Getenv("WS_HOST")
		srv.WsPath = os.Getenv("WS_PATH")
	}

	var persistence = func(srv *server.Server) {
		srv.ConnectionString = os.Getenv("CONNECTION_STRING")
	}

	http := server.New(app, persistence)
	http.Start()
}
