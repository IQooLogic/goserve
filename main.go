package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// ./goserve -h 0.0.0.0 -p 8080 -s -c /home/milos/go/src/github.com/baithive/certs/server.crt -k /home/milos/go/src/github.com/baithive/certs/server.key -d /home/milos/NetBeansProjects/ast-website/
func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nShutting down ...")
		os.Exit(0)
	}()

	host := flag.String("h", "0.0.0.0", "host")
	port := flag.Int("p", 8080, "port")
	tls := flag.Bool("s", false, "server over tls")
	cert := flag.String("c", "server.crt", "certificate")
	key := flag.String("k", "server.key", "private key")
	dir := flag.String("d", ".", "directory to serve")
	flag.Parse()

	var addr = *host + ":" + strconv.Itoa(*port)
	log.Printf("Serving %s on %s\n", *dir, addr)
	server := http.NewServeMux()
	fs := http.FileServer(http.Dir(*dir))
	server.Handle("/", fs)
	if *tls {
		err := http.ListenAndServeTLS(addr, *cert, *key, server)
		if err != nil {
			log.Fatalf("Error TLS: %s", err)
		}
	} else {
		err := http.ListenAndServe(addr, server)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
}
