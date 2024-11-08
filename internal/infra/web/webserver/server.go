package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WebServer struct {
	Router        *mux.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        mux.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (ws *WebServer) AddHandler(name string, handler http.HandlerFunc) {
	ws.Handlers[name] = handler
}

func (ws *WebServer) RegisterHandlers() {
	ws.Router.PathPrefix("/docs").Handler(ws.Handlers["docs"]).Methods("GET")
	ws.Router.HandleFunc("/products", ws.Handlers["createProduct"]).Methods("POST")
	ws.Router.HandleFunc("/products", ws.Handlers["listProducts"]).Methods("GET")
	ws.Router.HandleFunc("/products/{id}", ws.Handlers["getProduct"]).Methods("GET")
	ws.Router.HandleFunc("/products/{id}", ws.Handlers["updateProduct"]).Methods("PUT")
	ws.Router.HandleFunc("/products/{id}", ws.Handlers["deleteProduct"]).Methods("DELETE")
}

func (ws *WebServer) Start() {
	fmt.Println("Web server running on port:", ws.WebServerPort)

	wsAddress := fmt.Sprintf(":%s", ws.WebServerPort)

	err := http.ListenAndServe(wsAddress, ws.Router)
	if err != nil {
		log.Fatal(err)
	}
}
