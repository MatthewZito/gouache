package session

// SessionProvider describes manages the storage of Sessions.
type SessionProvider interface {
	// NewSession, given a session identifier, initializes a new Session.
	NewSession(sid string) (Session, error)
	// ReadSession, given a session identifier, returns the corresponding Session.
	ReadSession(sid string) (Session, error)
	// DestroySession, given a session identifier, destroys the corresponding Session.
	DestroySession(sid string) error
	// UpdateSession updates the Session's lastAccessed state. This method should handle any cache updates e.g. LRU.
	UpdateSession(sid string) error
	// FinalizeSessions garbage collects destroyed Sessions.
	// This method will only be invoked if the SessionManager `shouldFinalize` flag is set.
	// If `shouldFinalize` is `false`, the implementation should be a noop.
	FinalizeSessions(ttl int64)
}

// RegisterProvider registers a new session Provider for use with a SessionManager.
// This method enforces idempotence and will return an error if given an existing SessionProvider.
func RegisterProvider(name string, provider SessionProvider) error {
	if provider == nil {
		return sessionError("the given SessionProvider is nil")
	}
	if _, dup := providers[name]; dup {
		return sessionError("the given SessionProvider %s has already been registered", name)
	}

	providers[name] = provider
	return nil
}
