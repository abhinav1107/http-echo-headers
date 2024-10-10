package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	hostFlag = flag.String("host", "0.0.0.0", "interface to listen on")
	portFlag = flag.String("port", "8080", "port to listen on")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	// health endpoint
	mux.HandleFunc("/health", httpHealth())

	// root endpoint
	mux.HandleFunc("/", httpHeaders())

	var address = strings.Join([]string{*hostFlag, *portFlag}, ":")

	server := &http.Server{Addr: address, Handler: mux}
	serverCh := make(chan struct{})
	go func() {
		log.Printf("[INFO] server is listening on %s\n", address)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[ERROR] server exited with: %s", err)
		}
		close(serverCh)
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	log.Printf("[INFO] received interrupt, shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("[ERR] failed to shutdown server: %s", err)
	}

	os.Exit(2)
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "ok"}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func httpHeaders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("[INFO] Request: %v %v %v %v\n", r.RemoteAddr, r.URL, r.Host, r.Header)

		response := make(map[string]interface{})

		hostName, ok := os.LookupEnv("HOSTNAME")
		if !ok {
			hostName = "NO_HOSTNAME_IN_ENVIRONMENT"
		}

		response["Hostname"] = hostName
		for key, values := range r.Header {
			response[key] = strings.Join(values, ", ")
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
