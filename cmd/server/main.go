package main

import (
	"log"

	"github.com/christianbergsoerensen2/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
