package session

import "time"

// Session represents an object that that reads and writes session values.
type Session interface {
	// Set sets a Session value.
	Set(key, value interface{}) error
	// Get retrieves a Session value.
	Get(key interface{}) interface{}
	// Delete deletes a Session value.
	Delete(key interface{}) error
	// SessionID retruns the Session ID.
	SessionID() string
}

// SessionState represents the internal state of a Session.
type SessionState struct {
	provider     *SessionProvider
	sid          string
	lastAccessed time.Time
	value        map[interface{}]interface{}
}

// Set sets a key/value pair on the SessionState.
func (store *SessionState) Set(key, value interface{}) error {
	store.value[key] = value
	provider.UpdateSession(store.sid)

	return nil
}

// Get retrieves a key's corresponding value from the SessionState.
func (store *SessionState) Get(key interface{}) interface{} {
	provider.UpdateSession(store.sid)
	if v, ok := store.value[key]; ok {
		return v
	}

	return nil
}

// Delete removes a key/value pair from the SessionState.
func (store *SessionState) Delete(key interface{}) error {
	delete(store.value, key)
	provider.UpdateSession(store.sid)

	return nil
}

// SessionID returns the SessionState identifier.
func (store *SessionState) SessionID() string {
	return store.sid
}
