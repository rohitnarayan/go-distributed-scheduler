package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go-distributed-scheduler/internal/pkg/config"
)

func Start() error {
	config.Init()
	handler := NewHandler()
	startServer(router(handler))
	return nil
}

func startServer(handler http.Handler) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(config.App.Server.Port),
		Handler:      handler,
		ReadTimeout:  config.App.Server.ReadTimeout,
		WriteTimeout: config.App.Server.WriteTimeout,
	}

	fmt.Println("starting server on port ", server.Addr)
	go listenAndServe(server)

	<-stop
	server.Shutdown(context.Background())
}

func listenAndServe(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic("failed to start the server : " + err.Error())
	}
}
