package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MatthewZito/gouache/cache"
	"github.com/MatthewZito/gouache/format"
	"github.com/google/uuid"
)

const cookieId = "gouache_session"

type SessionContext struct {
	cache cache.SerializableStore
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewSessionContext(client cache.SerializableStore) SessionContext {
	return SessionContext{cache: client}
}

func (ctx SessionContext) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieId)

		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			token := cookie.Value
			session, err := ctx.cache.Get(token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

			} else if session.IsExpired() {
				ctx.cache.Delete(token)
				w.WriteHeader(http.StatusUnauthorized)
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (ctx SessionContext) Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
	}

	var expectedPassword = "password"
	if credentials.Password != expectedPassword {
		format.FormatError(w, http.StatusUnauthorized, "invalid credentials")
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	session := cache.Session{
		Username: credentials.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.cache.Set(sessionToken, session); err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    sessionToken,
		Expires:  time.Now().Add(120 * time.Second),
		HttpOnly: true,
		Path:     "/",
	})

	format.FormatResponse(w, http.StatusOK, nil)
}

func (ctx SessionContext) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(cookieId)
	sessionToken := cookie.Value

	ctx.cache.Delete(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})

	format.FormatResponse(w, http.StatusOK, nil)
}

func (ctx SessionContext) RenewSession(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(cookieId)
	sessionToken := cookie.Value

	session, err := ctx.cache.Get(sessionToken)
	if err != nil {
		format.FormatError(w, http.StatusUnauthorized, err.Error())
	}

	expiresAt := time.Now().Add(120 * time.Second)

	session.Expiry = expiresAt

	ctx.cache.Set(sessionToken, *session)

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})
}
