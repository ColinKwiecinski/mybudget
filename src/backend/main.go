// Placeholder for main file
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

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")
	if tlsKeyPath == "" || tlsCertPath == "" {
		log.Fatal("Environrment not set")
	}
	mux := mux.NewRouter()

	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}
