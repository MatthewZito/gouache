package session

import (
	"container/list"
	"sync"
	"time"
)

// MemoryProvider represents an object that provides persistent storage for Sessions.
type MemoryProvider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	state    *list.List
}

func NewMemoryProvider() *MemoryProvider {
	return &MemoryProvider{state: list.New()}
}

func (provider *MemoryProvider) NewSession(sid string) (Session, error) {
	provider.lock.Lock()
	defer provider.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)

	session := Session{provider: provider, sid: sid, lastAccessed: time.Now(), value: v}
	node := provider.state.PushBack(session)
	provider.sessions[sid] = node

	return session, nil
}

func (provider *MemoryProvider) ReadSession(sid string) (Session, error) {
	if node, ok := provider.sessions[sid]; ok {
		return node.Value.(Session), nil
	} else {
		session, err := provider.NewSession(sid)
		return session, err
	}
}

func (provider *MemoryProvider) DestroySession(sid string) error {
	if node, ok := provider.sessions[sid]; ok {
		delete(provider.sessions, sid)
		provider.state.Remove(node)
	}

	return nil
}

func (provider *MemoryProvider) FinalizeSessions(ttl int64) {
	provider.lock.Lock()
	defer provider.lock.Unlock()

	for {
		node := provider.state.Back()
		if node == nil {
			break
		}

		if node.Value.(*Session).lastAccessed.Unix()+ttl < time.Now().Unix() {
			provider.DestroySession(node.Value.(*Session).sid)
		} else {
			break
		}
	}
}

func (provider *MemoryProvider) UpdateSession(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()

	if node, ok := provider.sessions[sid]; ok {
		node.Value.(*Session).lastAccessed = time.Now()
		provider.state.MoveToFront(node)
	}

	return nil
}
