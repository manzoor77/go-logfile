package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create a log file
	logFile, err := os.OpenFile("api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Create a logger that writes to the log file
	logger := log.New(logFile, "", log.LstdFlags)

	// Define an HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)
		fmt.Fprintln(w, "Hello, World!")
	})

	// Create an HTTP server with the handler
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the HTTP server
	logger.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
