// @todo convert to internal pkg
package session

import (
	"time"
)

type MockProvider struct {
	sessions map[string]*Session
}

func NewMockProvider() *MockProvider {
	return &MockProvider{sessions: make(map[string]*Session)}
}

func (m *MockProvider) NewSession(sid string) (Session, error) {
	sess := Session{provider: m, sid: sid, lastAccessed: time.Now(), value: map[interface{}]interface{}{"id": sid}}

	m.sessions[sid] = &sess
	return sess, nil
}

func (m *MockProvider) ReadSession(sid string) (Session, error) {
	return *m.sessions[sid], nil
}

func (m *MockProvider) DestroySession(sid string) error {
	delete(m.sessions, sid)
	return nil
}

func (m *MockProvider) UpdateSession(sid string) error {
	return nil
}

func (m *MockProvider) FinalizeSessions(ttl int64) {}
