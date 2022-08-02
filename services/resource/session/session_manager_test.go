package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewSessionManager(t *testing.T) {
	t.Run("InvalidProviderError", func(t *testing.T) {
		s, err := NewSessionManager("invalid", "", 0, false)
		if err == nil {
			t.Errorf("expected err to be an error but got nil\n")
		}

		if s != nil {
			t.Errorf("expected session to be an nil but got %v\n", s)
		}
	})

	t.Run("NewSessionManager", func(t *testing.T) {
		RegisterProvider("test", NewMockProvider())

		sm, err := NewSessionManager("test", "test", int64(time.Now().Add(60*time.Minute).Second()), false)

		if err != nil {
			t.Errorf("expected err to be nil but got %v\n", err)
		}

		if sm == nil {
			t.Errorf("expected SessionManager to be initialized but got nil\n")
		}
	})
}

func TestSetCookie(t *testing.T) {
	RegisterProvider("test", NewMockProvider())

	sm, _ := NewSessionManager("test", "test", int64(time.Until(time.Now().Add(60*time.Minute)).Seconds()), false)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	cookieBefore := rec.Header()["Set-Cookie"]
	if len(cookieBefore) != 0 {
		t.Errorf("expected Set-Cookie header to be omitted but got %v\n", cookieBefore)
	}

	_, err := sm.NewSession(rec, req)
	if err != nil {
		t.Errorf("did not expect an error but got %v\n", err)
	}

	cookieAfter := rec.Header()["Set-Cookie"]
	if len(cookieAfter) != 1 {
		t.Errorf("expected Set-Cookie header to be set but it wasn't\n")
	}
}

func TestDestroySession(t *testing.T) {
	cookieName := "test"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	p := NewMockProvider()
	RegisterProvider("test", p)

	sm, _ := NewSessionManager("test", cookieName, int64(time.Until(time.Now().Add(60*time.Minute)).Seconds()), false)
	sm.NewSession(rec, req)

	req2 := &http.Request{Header: http.Header{"Cookie": rec.Header()["Set-Cookie"]}}

	if err := sm.DestroySession(rec, req2); err != nil {
		t.Errorf("expected DestroySession to succeed but got err %v\n", err)
	}

	cookie, err := req2.Cookie(cookieName)
	if err != nil {
		t.Errorf("expected cookie to be set on request but it was %v\n", cookie)
	}

	if cookie.Name != cookieName {
		t.Errorf("expected cookie name to be %s but got %s\n", cookieName, cookie.Name)
	}

	if cookie.MaxAge != 0 {
		t.Errorf("expected cookie MaxAge to be 0 but got %v\n", cookie.MaxAge)
	}
}

func TestBuildCookie(t *testing.T) {
	cookieName := "test_cookie_name"
	sid := "123"

	RegisterProvider("test", NewMockProvider())

	sm, _ := NewSessionManager("test", cookieName, int64(time.Until(time.Now().Add(60*time.Minute)).Seconds()), false)
	cookie := sm.buildCookie(sid, int(sm.ttl))

	// @todo make this an option
	if cookie.HttpOnly != true {
		t.Errorf("expected HttpOnly to be true but got %v\n", cookie.HttpOnly)
	}

	if cookie.Name != cookieName {
		t.Errorf("expected cookie name to be %s but got %s\n", cookieName, cookie.Name)
	}

	if cookie.MaxAge != 3599 {
		t.Errorf("expected cookie MaxAge to be 3599 but got %v\n", cookie.MaxAge)
	}

	if cookie.Value != sid {
		t.Errorf("expected cookie value to be %s but got %s\n", sid, cookie.Value)
	}
}
