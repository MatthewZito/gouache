package services

import (
	"encoding/json"
	"time"

	"github.com/exbotanical/gouache/entities"
	"github.com/exbotanical/gouache/repositories"
)

// SessionService defines the contract for a `SessionServiceProvider` instance.
type SessionService interface {
	GetSession(string) (*entities.Session, error)
	SetSession(string, entities.Session) error
	DeleteSession(string) error
}

// SessionServiceProvider is a service layer API for interacting with session data.
type SessionServiceProvider struct {
	*repositories.SessionRepository
}

// Create a new SessionService.
func NewSessionService() (SessionService, error) {
	r, err := repositories.NewSessionRepository()
	if err != nil {
		return nil, err
	}

	return &SessionServiceProvider{
		r,
	}, nil
}

// DeleteSession - when provided a session ID `sid` - deletes the corresponding `Session` from the cache.
func (s *SessionServiceProvider) DeleteSession(sid string) error {
	if _, err := s.Client.Del(s.Ctx, sid).Result(); err != nil {
		return err
	}

	return nil
}

// GetSession - when provided a session ID `sid` - retrieves the corresponding `Session` from the cache.
func (s *SessionServiceProvider) GetSession(sid string) (*entities.Session, error) {
	if v, err := s.Client.Get(s.Ctx, sid).Bytes(); err != nil {
		return nil, err
	} else {
		session := &entities.Session{}

		if err := json.Unmarshal(v, session); err != nil {
			return nil, err
		}

		return session, nil
	}
}

// SetSession - when provided a session ID `sid` and `Session` `session` - sets or updates the `Session` in the cache.
func (s *SessionServiceProvider) SetSession(sid string, session entities.Session) error {
	v, err := json.Marshal(session)
	if err != nil {
		return err
	}

	if err := s.Client.Set(s.Ctx, sid, v, time.Until(session.Expiry)).Err(); err != nil {
		return err
	}

	return nil
}
