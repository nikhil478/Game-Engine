package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/nikhil478/api_server/engine"
)

func main() {
	port := flag.String("port", "8080", "API server port")
	flag.StringVar(&engine.EngineURL, "engine-url", "http://localhost:8081/evaluate", "Game Engine URL")
	flag.Parse()
	// TODO: input sanitization

	http.HandleFunc("/submit", engine.SubmitHandler)

	log.Printf("API Server running on port %s, forwarding to engine %s", *port, engine.EngineURL)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
	// TODO: graceful shutdown
}
