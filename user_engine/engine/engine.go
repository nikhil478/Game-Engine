package engine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/nikhil478/user_engine/engine/models"
)

func MockUserEngine(numUsers int) {
	var wg sync.WaitGroup
	wg.Add(numUsers)

	for i := 1; i <= numUsers; i++ {
		go func(userID int) {
			defer wg.Done()

			delay := rand.Intn(991) + 10
			time.Sleep(time.Duration(delay) * time.Millisecond)

			answer := "no"
			if rand.Intn(2) == 1 {
				answer = "yes"
			}

			resp := models.UserResponse{
				UserID:     userID,
				Answer:     answer,
				DelayMS:    delay,
				Successful: true,
			}

			err := sendResponse(resp)
			if err != nil {
				fmt.Printf("User %d failed: %v\n", userID, err)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Simulation completed.")
}
