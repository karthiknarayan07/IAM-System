package http

import (
	"log"
	"net/http"
)

func StartServer(router *Router, port string) {
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, router.Engine); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
