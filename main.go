package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RisingStack/almandite-user-service/handlers"
	"github.com/RisingStack/almandite-user-service/middleware"
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
		log.Fatal(err)
	}

	tcpAddr := listener.Addr().(*net.TCPAddr)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-signalCh:
			log.Println("Signal received, shutting down...")
			CloseDBConnection()
			os.Exit(0)
		}
	}()

	OpenDBConnection()
	if err := Migrate(); err != nil {
		log.Fatal(err)
	}

	users, err := UserRepository.Fetch()
	if err != nil {
		log.Println("Failed to get users", err)
	}
	log.Println("Users", users)

	log.Printf("Open the following URL in the browser: http://%s:%d\n", convertIPtoString(tcpAddr.IP), tcpAddr.Port)

	http.HandleFunc("/healthcheck",
		middleware.Chain(
			middleware.Timer,
			middleware.Logger,
		)(handlers.Healthcheck))

	if err := http.Serve(listener, nil); err != nil {
		log.Fatal(err)
	}
}

func convertIPtoString(ip net.IP) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
