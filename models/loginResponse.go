package models

type LoginResponse struct {
	// omitempty = is error then return struct empty
	Token string `json:"token,omitempty"`
}
