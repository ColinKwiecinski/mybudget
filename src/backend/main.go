package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80" // TODO: Change to 443 When implementing TLS
	}

	// TODO: update this to TLS mux
	mux := http.NewServeMux()

	log.Printf("Listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func getEnv(name string) string{
	result := os.Getenv(name)
	if (len(result) == 0) {
		envNotFound(name)
	}
	return result
}

func envNotFound(name string) {
	log.Fatalf("%s not set or not found", name)
}
