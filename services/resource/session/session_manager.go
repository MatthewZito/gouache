package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// SessionManager represents an object that manages all Sessions in the service.
type SessionManager struct {
	cookieName string
	lock       sync.Mutex
	provider   *SessionProvider
	ttl        int64
}

// NewSessionManager initializes and returns a new SessionManager object.
func NewSessionManager(providerName string, cookieName string, ttl int64) (*SessionManager, error) {
	if provider, ok := providers[providerName]; !ok {
		return nil, sessionError("unknown provider %q", providerName)
	} else {
		return &SessionManager{provider: provider, cookieName: cookieName, ttl: ttl}, nil
	}
}

// NewSessionId creates a new pseudo-unique Session identifier.
func (manager *SessionManager) NewSessionId() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", sessionError("error reading bytes %v", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// NewSession creates a new Session object.
func (manager *SessionManager) NewSession(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// Attempt to retrieve the Cookie.
	cookie, err := r.Cookie(manager.cookieName)

	// If the Cookie does not exist, or has no value, create a new one.
	if err != nil || cookie.Value == "" {
		// Generate a new SessionId.
		sid, err := manager.NewSessionId()
		if err != nil {
			sessionError("error creating new SessionId %v", err)
		}

		// Create a new Session.
		session, _ = manager.provider.NewSession(sid)
		// Create the new Cookie.
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.ttl)}
		// Set the Cookie.
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.ReadSession(sid)
	}

	return
}

// DestroySession destroys the current Session and updates the corresponding Cookie.
func (manager *SessionManager) DestroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)

	if err == nil && cookie.Value != "" {
		manager.lock.Lock()
		defer manager.lock.Unlock()

		manager.provider.DestroySession(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *SessionManager) FinalizeSessions() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.provider.FinalizeSessions(manager.ttl)
	time.AfterFunc(time.Duration(manager.ttl), func() { manager.FinalizeSessions() })
}
