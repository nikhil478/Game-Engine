package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Event struct {
	Response UserResponse
	Result   chan struct{}
}

type UserResponse struct {
	UserID     int    `json:"user_id"`   
	Answer     string `json:"answer"`   
	DelayMS    int    `json:"delay_ms"`  
	Successful bool   `json:"successful"`
}

func main() {
	port := flag.String("port", "9090", "Game Engine port")
	flag.Parse()

	events := make(chan Event, 1000)
	var winnerSet bool
	var winnerID int
	var correctCount, incorrectCount int
	var mu sync.Mutex
	startTime := time.Now()

	go func() {
		for ev := range events {
			resp := ev.Response
			mu.Lock()
			if resp.Successful && resp.Answer == "yes" {
				correctCount++
				if !winnerSet {
					winnerSet = true
					winnerID = resp.UserID
					fmt.Printf("🏆 WINNER: user_id=%d\n", winnerID)
					fmt.Printf("Time taken to find winner: %v\n", time.Since(startTime))
				}
			} else {
				incorrectCount++
			}
			mu.Unlock()
			ev.Result <- struct{}{}
		}
	}()

	http.HandleFunc("/evaluate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var ur UserResponse
		if err := json.NewDecoder(r.Body).Decode(&ur); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		ev := Event{
			Response: ur,
			Result:   make(chan struct{}),
		}

		events <- ev
		<-ev.Result

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"processed": true,
			"winner_id": winnerID,
		})
	})

	log.Printf("Game Engine listening on :%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
