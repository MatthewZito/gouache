package entities

import "time"

// Session represents session metadata as stored in the redis session cache.
type Session struct {
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

// SessionResponse represents session metadata provided to the client.
type SessionResponse struct {
	Username string `json:"username"`
	Exp      int    `json:"exp"`
}

// IsExpired is a getter that determines whether the `Session` is expired.
func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
