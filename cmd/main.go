package main

import (
	"flag"

	"sample/pkg/presentation/web"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	web.Serve()
}
