package entities

// SessionResponse represents session metadata provided to the client.
type SessionResponse struct {
	Username string `json:"username"`
	Exp      int    `json:"exp"`
}
