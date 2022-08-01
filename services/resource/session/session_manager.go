package session

import (
	"net/http"
	"sync"
	"time"
)

// SessionManager represents an object that manages all Sessions in the service.
type SessionManager struct {
	cookieName     string
	lock           sync.Mutex
	provider       SessionProvider
	ttl            int64
	shouldFinalize bool
}

// NewSessionManager initializes and returns a new SessionManager object.
func NewSessionManager(providerName string, cookieName string, ttl int64, shouldFinalize bool) (*SessionManager, error) {
	if provider, ok := providers[providerName]; !ok {
		return nil, sessionError("unknown provider %s", providerName)
	} else {
		return &SessionManager{
			provider:       provider,
			cookieName:     cookieName,
			ttl:            ttl,
			shouldFinalize: shouldFinalize,
		}, nil
	}
}

// NewSession creates a new Session object.
func (manager *SessionManager) NewSession(w http.ResponseWriter, r *http.Request) (*Session, error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// Attempt to retrieve the Cookie.
	cookie, err := r.Cookie(manager.cookieName)

	// If the Cookie does not exist, or has no value, create a new one.
	if err != nil || cookie.Value == "" {
		// Generate a new SessionId.
		sid := newSessionId()

		// Create a new Session.
		session, err := manager.provider.NewSession(sid)
		if err != nil {
			return nil, sessionError("error creating new Session %v", err)
		}

		// Create the new Cookie.
		cookie := manager.buildCookie(sid)
		// Set the Cookie.
		http.SetCookie(w, &cookie)

		return &session, nil
	}

	session, err := manager.provider.ReadSession(cookie.Value)

	if err != nil {
		return nil, sessionError("error reading existing Session %v", err)
	}

	return &session, nil
}

// DestroySession destroys the current Session and updates the corresponding Cookie.
func (manager *SessionManager) DestroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)

	if err == nil && cookie.Value != "" {
		manager.lock.Lock()
		defer manager.lock.Unlock()

		manager.provider.DestroySession(cookie.Value)

		expiration := time.Now()
		manager.ttl = -1
		cookie := manager.buildCookie(cookie.Value)
		cookie.Expires = expiration

		http.SetCookie(w, &cookie)
	}
}

func (manager *SessionManager) FinalizeSessions() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.provider.FinalizeSessions(manager.ttl)
	time.AfterFunc(time.Duration(manager.ttl), func() { manager.FinalizeSessions() })
}

func (manager *SessionManager) buildCookie(sid string) http.Cookie {
	return http.Cookie{
		Name:     manager.cookieName,
		Value:    sid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(manager.ttl),
	}
}
