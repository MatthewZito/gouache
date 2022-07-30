package session

import (
	"container/list"
	"sync"
	"time"
)

// SessionProvider represents an object that provides persistent storage for Sessions.
type SessionProvider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	state    *list.List
}

// SessionProviderActor describes SessionProvider methods that manage reads and writes to Sessions.
type SessionProviderActor interface {
	// NewSession, given a session identifier, initializes a new Session.
	NewSession(sid string) (Session, error)
	// ReadSession, given a session identifier, returns the corresponding Session.
	ReadSession(sid string) (Session, error)
	// DestroySession, given a session identifier, destroys the corresponding Session.
	DestroySession(sid string) error
	// FinalizeSessions garbage collects destroyed Sessions.
	FinalizeSessions(ttl int64)
}

// RegisterProvider registers a new session Provider for use with a SessionManager.
// This method enforces idempotence and will return an error if given an existing SessionProvider.
func RegisterProvider(name string, provider *SessionProvider) error {
	if provider == nil {
		return sessionError("the given SessionProvider is nil")
	}

	if _, dup := providers[name]; dup {
		return sessionError("the given SessionProvider %s has already been registered", name)
	}

	providers[name] = provider
	return nil
}

func (provider *SessionProvider) NewSession(sid string) (Session, error) {
	provider.lock.Lock()
	defer provider.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)

	session := &SessionState{sid: sid, lastAccessed: time.Now(), value: v, provider: provider}
	node := provider.state.PushBack(session)
	provider.sessions[sid] = node

	return session, nil
}

func (provider *SessionProvider) ReadSession(sid string) (Session, error) {
	if node, ok := provider.sessions[sid]; ok {
		return node.Value.(*SessionState), nil
	} else {
		session, err := provider.NewSession(sid)
		return session, err
	}
}

func (provider *SessionProvider) DestroySession(sid string) error {
	if node, ok := provider.sessions[sid]; ok {
		delete(provider.sessions, sid)
		provider.state.Remove(node)
	}

	return nil
}

func (provider *SessionProvider) FinalizeSessions(ttl int64) {
	provider.lock.Lock()
	defer provider.lock.Unlock()

	for {
		node := provider.state.Back()
		if node == nil {
			break
		}

		if node.Value.(*SessionState).lastAccessed.Unix()+ttl < time.Now().Unix() {
			provider.DestroySession(node.Value.(*SessionState).sid)
		} else {
			break
		}
	}
}

func (provider *SessionProvider) UpdateSession(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()

	if node, ok := provider.sessions[sid]; ok {
		node.Value.(*SessionState).lastAccessed = time.Now()
		provider.state.MoveToFront(node)
	}

	return nil
}
