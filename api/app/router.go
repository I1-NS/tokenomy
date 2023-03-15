// Copyright 2023 © Tokenomy. All rights reserved.
package app

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/i1-ns/tokenomy/api/handler"
	"github.com/i1-ns/tokenomy/api/middleware"
)

func New(addr string) *http.ServeMux {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	router := http.DefaultServeMux
	server := &http.Server{
		Addr:           addr,
		Handler: 		middleware.Request(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { fmt.Printf("listen: %s\n", err.Error()) }
	}()
	router.HandleFunc("/api/v0/tokenomy", handler.GetByIDs)
	fmt.Printf("Listening on port %s\n", addr)
	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	println("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 5 seconds to finish, the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil { fmt.Printf("Server forced to shutdown: ", err) }
	println("Server exiting")
	return router
}