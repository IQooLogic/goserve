package action

import (
	"github.com/iqoologic/goserve/server"
	"github.com/urfave/cli/v2"
)

func Run(context *cli.Context) error {
	host := context.String("u")
	port := context.Int("p")
	dir := context.String("d")
	tls := context.Bool("s")
	cert := context.String("c")
	key := context.String("k")

	srv := server.New()
	if tls {
		srv.StartTLS(host, port, dir, cert, key)
	} else {
		srv.Start(host, port, dir)
	}

	return nil
}
