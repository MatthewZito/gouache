package session

import "time"

// SessionState manages the internal state of a Session.
type SessionState interface {
	// Set sets a Session value.
	Set(key, value interface{}) error
	// Get retrieves a Session value.
	Get(key interface{}) interface{}
	// Delete deletes a Session value.
	Delete(key interface{}) error
	// SessionID returns the Session ID.
	SessionID() string
}

// Session stores session state.
type Session struct {
	provider     SessionProvider
	sid          string
	lastAccessed time.Time
	value        map[interface{}]interface{}
}

// Set sets a key/value pair on the SessionState and invokes the SessionProvider's UpdateSession method.
func (store *Session) Set(key, value interface{}) error {
	store.value[key] = value
	store.provider.UpdateSession(store.sid)

	return nil
}

// Get retrieves a key's corresponding value from the Session and invokes the SessionProvider's UpdateSession method.
func (store *Session) Get(key interface{}) interface{} {
	store.provider.UpdateSession(store.sid)
	if v, ok := store.value[key]; ok {
		return v
	}

	return nil
}

// Delete removes a key/value pair from the Session and invokes the SessionProvider's UpdateSession method.
func (store *Session) Delete(key interface{}) error {
	delete(store.value, key)
	store.provider.UpdateSession(store.sid)

	return nil
}

// SessionID returns the Session identifier.
func (store *Session) SessionID() string {
	return store.sid
}
