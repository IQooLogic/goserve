package server

import (
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func addHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s - %s -> %s", r.Method, r.URL.Path, r.RemoteAddr, r.Host)
		w.Header().Add("X-Frame-Options", "DENY")
		fs.ServeHTTP(w, r)
	}
}

func RandomizePort() int {
	rand.Seed(time.Now().UnixNano())
	rndPort := rand.Intn(9999-8000+1) + 8000
	return rndPort
}

func addresses() string {
	addresses, _ := net.InterfaceAddrs()
	length := len(addresses)
	ret := make([]string, length)
	for i, address := range addresses {
		ret[i] = strings.Split(address.String(), "/")[0]
	}

	return strings.Join(ret, ", ")
}

func CurrentDirectory() string {
	path, _ := os.Getwd()
	return path
}
