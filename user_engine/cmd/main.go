package main

import (
	"flag"
	"fmt"

	"github.com/nikhil478/user_engine/engine"
)

func main() {

	numUsers := flag.Int("n", 1000, "number of simulated users")
	api := flag.String("url", "http://localhost:8080/api", "API server URL")
	flag.Parse()

	engine.ApiURL = *api

	fmt.Printf("Simulating %d users hitting %s\n", *numUsers, engine.ApiURL)

	engine.MockUserEngine(*numUsers)
}
