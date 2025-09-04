package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikhil478/user_engine/engine/models"
)

var ApiURL string

func sendResponse(resp models.UserResponse) error {
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	r, err := http.Post(ApiURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with %d", r.StatusCode)
	}

	return nil
}
