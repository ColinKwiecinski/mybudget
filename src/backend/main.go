package main

import (
	"log"
	"net/http"
	"os"
<<<<<<< HEAD

	"github.com/gorilla/mux"
=======
>>>>>>> 07c50e2feae86c7c5640e00f51ef0ffb64a23de3
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
<<<<<<< HEAD
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")
	if tlsKeyPath == "" || tlsCertPath == "" {
		log.Fatal("Environrment not set")
	}
	mux := mux.NewRouter()

	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
=======
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
>>>>>>> 07c50e2feae86c7c5640e00f51ef0ffb64a23de3
}
