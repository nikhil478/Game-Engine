package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nikhil478/api_server/engine/models"
)

// TODO : remove this func
func MockSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var resp models.UserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("request body : %v \n", resp)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var resp models.UserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	engineResp, err := ForwardToEngine(resp)
	if err != nil {
		log.Printf("Error forwarding to engine: %v", err)
		http.Error(w, "failed to contact game engine", http.StatusBadGateway)
		return
	}
	defer engineResp.Body.Close()

	body, _ := io.ReadAll(engineResp.Body)

	var jsonResp interface{}
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		fmt.Printf("response body (raw): %s\n", string(body))
	} else {
		prettyJSON, _ := json.MarshalIndent(jsonResp, "", "  ")
		fmt.Printf("response body:\n%s\n", string(prettyJSON))
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(engineResp.StatusCode)
	w.Write(body)
}
