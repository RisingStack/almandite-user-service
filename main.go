package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

const DefaultHTTPAddr = ":0"

var httpAddr string

func init() {
	flag.StringVar(&httpAddr, "addr", DefaultHTTPAddr, "Set the HTTP bind address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		panic(err)
	}

	tcpAddr := listener.Addr().(*net.TCPAddr)

	fmt.Printf("Open the following URL in the browser: http://%s:%d", convertIPtoString(tcpAddr.IP), tcpAddr.Port)

	http.Serve(listener, nil)
}

func convertIPtoString(ip net.IP) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
