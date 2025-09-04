package models

type UserResponse struct {
	UserID     int    `json:"user_id"`
	Answer     string `json:"answer"`
	DelayMS    int    `json:"delay_ms"`
	Successful bool   `json:"successful"`
}