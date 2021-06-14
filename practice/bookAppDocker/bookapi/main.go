package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/book-api/utilities"
)

var myEnv map[string]string = utilities.LoadEnvVaribles()
var shutdownTimeout = flag.Duration("shutdown-timeout", 10*time.Second, "shutdown timeout (5s,5m,5h) before connections are cancelled")

const (
	service = "book-service"
)

func main() {
	r := mux.NewRouter()

	RegisterHandlers(r)

	serverIp := myEnv["SERVER_IP"]
	serverPort := myEnv["SERVER_PORT"]

	srv := &http.Server{
		Addr:    serverIp + ":" + serverPort,
		Handler: r,
	}

	gracefulShutdown(srv, serverIp, serverPort)

}

func gracefulShutdown(srv *http.Server, serverIp string, serverPort string) {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("%s listening on %s:%s with %v timeout", service, serverIp, serverPort, *shutdownTimeout)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stop

	log.Printf("%s shutting down ...\n", service)

	ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s down\n", service)
}
