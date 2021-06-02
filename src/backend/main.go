package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := getEnv("TLSKEY")
	tlsCertPath := getEnv("TLSCERT")

	mux := mux.NewRouter()

	log.Printf("Listening at %s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}

func getEnv(name string) string {
	result := os.Getenv(name)
	if len(result) == 0 {
		envNotFound(name)
	}
	return result
}

func envNotFound(name string) {
	log.Fatalf("%s not set or not found", name)
}
