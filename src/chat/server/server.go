package server

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	handlers "prismapp-test/src/chat/handlers"
	retrieveUseCase "prismapp-test/src/chat/usecase/retrieve"
	sendUseCase "prismapp-test/src/chat/usecase/send"
	gateway "prismapp-test/src/chat/utility/mysql"
	wsocket "prismapp-test/src/chat/utility/websocket"
	"sync"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
)

type HTTPServers interface {
	Start()
}

type Server struct {

	// APP
	Address string
	WsHost  string
	WsPath  string

	// Persistence
	ConnectionString string
}

type service struct {
	server *Server
}

// New constructor
func New(options ...func(*Server)) HTTPServers {

	s := &Server{}

	// call option functions on instance to set options on it
	for _, opt := range options {
		opt(s)
	}

	return &service{
		server: s,
	}
}

// templ represents a single template
type templateHandler struct {

	// sync.Once menjamin jika fungsi yang akan dilewati hanya diesekusi sekali saja
	// tidak peduli seberapa sering goroutines memanggil method ServerHTTP
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("../../views", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func (s *service) Start() {

	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stderr)
		logger = &serializedLogger{Logger: logger}
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	}

	connection := gateway.Connection(s.server.ConnectionString)
	roomsGateway := gateway.NewRoomGateway(connection)

	wsConnect := wsocket.Connection(s.server.WsHost, s.server.WsPath)
	websocketGateway := wsocket.NewPublishGateway(wsConnect)

	mux := http.NewServeMux()

	var snmService sendUseCase.Service
	snmService = sendUseCase.NewService(roomsGateway, websocketGateway)
	snmService = handlers.NewLoggingMiddleware(kitlog.With(logger, "service", "SendNewMessage"), snmService)

	var rapmService retrieveUseCase.Service
	rapmService = retrieveUseCase.NewService(roomsGateway)

	mux.Handle("/chats", handlers.MakeHandler(snmService, rapmService, logger))
	mux.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/", accessControl(mux))

	var addr = flag.String("addr", ":"+s.server.Address, "The addr of the  application.")

	log.Println("Starting web server on", *addr)
	errs := make(chan error, 2)

	go func() {
		logger.Log("transport", "http", "address", *addr, "msg", "listening")
		errs <- http.ListenAndServe(*addr, nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

type serializedLogger struct {
	mtx sync.Mutex
	kitlog.Logger
}
