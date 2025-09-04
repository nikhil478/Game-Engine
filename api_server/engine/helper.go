package engine

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/nikhil478/api_server/engine/models"
)

var EngineURL string

func ForwardToEngine(resp models.UserResponse) (*http.Response, error) {
	payload, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", EngineURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
