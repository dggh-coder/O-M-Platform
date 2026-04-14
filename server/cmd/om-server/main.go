package main

import (
	"log"
	"net/http"
	"os"

	"om-platform/server/internal/api"
)

func main() {
	mux := api.NewRouter()
	addr := os.Getenv("OM_SERVER_ADDR")
	if addr == "" {
		addr = ":8081"
	}

	log.Printf("O&M server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
