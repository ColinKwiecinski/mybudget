package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"mybudget.com/src/backend/auth"
	"mybudget.com/src/backend/sessions"
)

type ResponseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(rh.headerName, rh.headerValue)
	rh.handler.ServeHTTP(w, r)
}

func NewResponseHeader(handler http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{handler, headerName, headerValue}
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := getEnv("TLSKEY")
	tlsCertPath := getEnv("TLSCERT")

	sessionKey := getEnv("SESSIONKEY")
	redisaddr := getEnv("REDISADDR")
	dsn := getEnv("DSN")
	hour, _ := time.ParseDuration("1h")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisaddr,
		Password: "",
		DB:       0,
	})
	userStore := auth.NewMysqlStore(dsn)

	context := auth.HandlerContext{
		SigningKey: sessionKey,
		Sessions:   sessions.NewRedisStore(redisClient, hour),
		Users:      userStore}

	mux := mux.NewRouter()

	mux.HandleFunc("/users", context.UsersHandler)
	mux.HandleFunc("/users/{UID}", context.SpecificUserHandler)
	mux.HandleFunc("/transactions", context.TransactionHandler)
	mux.HandleFunc("/transactions/{UID}", context.SpecificTransactionHandler)

	wrappedMux1 := NewResponseHeader(mux, "Access-Control-Allow-Origin", "*")
	wrappedMux2 := NewResponseHeader(wrappedMux1, "Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	wrappedMux3 := NewResponseHeader(wrappedMux2, "Access-Control-Allow-Headers", "Content-Type, Authorization")
	wrappedMux4 := NewResponseHeader(wrappedMux3, "Access-Control-Expose-Headers", "Authorization")
	wrappedMux5 := NewResponseHeader(wrappedMux4, "Access-Control-Max-Age", "600")

	log.Printf("Listening at %s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux5))
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
