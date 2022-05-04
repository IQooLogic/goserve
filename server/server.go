package server

import (
	"fmt"
	"github.com/iqoologic/goserve/server2"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (srv *Server) Start(host string, port int, dir string) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving '%s' on '%s'", dir, addr)
	if host == "0.0.0.0" {
		log.Printf("Listening on: %s", addresses())
	}

	server := http.NewServeMux()
	//fs := http.FileServer(http.Dir(dir))
	fs2 := server2.MyFileServer(http.Dir(dir))
	server.Handle("/", addHeaders(fs2))
	go func() {
		err := http.ListenAndServe(addr, server)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}()
}

func (srv *Server) StartTLS(host string, port int, dir, cert, key string) {
	var addr = host + ":" + strconv.Itoa(port)
	log.Printf("Serving '%s' on '%s'", dir, addr)
	if host == "0.0.0.0" {
		log.Printf("Listening on: %s", addresses())
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/", addHeaders(fs))
	go func() {
		err := http.ListenAndServeTLS(addr, cert, key, mux)
		if err != nil {
			log.Fatalf("Error TLS: %s", err)
		}
	}()
}
