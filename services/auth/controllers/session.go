package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/exbotanical/gouache/cache"
	"github.com/exbotanical/gouache/models"
	"github.com/exbotanical/gouache/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const cookieId = "gouache_session"

type SessionContext struct {
	cache cache.SerializableStore
	repo  *repositories.DB
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewSessionContext(client cache.SerializableStore, repo *repositories.DB) SessionContext {
	return SessionContext{cache: client, repo: repo}
}

func (ctx SessionContext) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieId)

		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			token := cookie.Value
			session, err := ctx.cache.Get(token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else if session.IsExpired() {
				ctx.cache.Delete(token)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (ctx SessionContext) Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "invalid credentials", 0)
		return
	}

	u, err := ctx.repo.GetUser(credentials.Username)
	if err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "invalid credentials", 0)
		return
	}

	if !CheckPasswordHash(credentials.Password, u.Password) {
		models.FormatError(w, http.StatusUnauthorized, "password mismatch", "invalid credentials", 0)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	session := cache.Session{
		Username: credentials.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.cache.Set(sessionToken, session); err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "an unknown exception occurred @todo const", 0)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	models.FormatResponse(w, http.StatusOK, models.DefaultOk(
		models.SessionResponse{
			Username: u.Username,
			Exp:      int(time.Until(session.Expiry).Seconds()),
		},
	),
	)
}

func (ctx SessionContext) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieId)
	if err != nil {
		models.FormatError(w, http.StatusUnauthorized, err.Error(), "an unknown exception occurred @todo const", 0)
		return
	}

	sessionToken := cookie.Value

	ctx.cache.Delete(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})

	models.FormatResponse(w, http.StatusOK, models.DefaultOk(nil))
}

func (ctx SessionContext) RenewSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieId)
	if err != nil {
		models.FormatError(w, http.StatusUnauthorized, err.Error(), "unauthorized", 0)
		return
	}

	sessionToken := cookie.Value

	session, err := ctx.cache.Get(sessionToken)
	if err != nil {
		models.FormatError(w, http.StatusUnauthorized, err.Error(), "unauthorized", 0)
		return
	}

	expiresAt := time.Now().Add(60 * time.Minute)
	session.Expiry = expiresAt

	ctx.cache.Set(sessionToken, *session)

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	models.FormatResponse(w, http.StatusOK, models.DefaultOk(models.SessionResponse{Username: session.Username, Exp: int(time.Until(session.Expiry).Seconds())}))
}

// @todo Determine when user exists
func (ctx SessionContext) Register(w http.ResponseWriter, r *http.Request) {
	u := &models.NewUserTemplate{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "invalid credentials provided", 0)
		return
	}

	hash, err := HashPassword(u.Password)
	if err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "invalid credentials provided", 0)
		return
	}

	u.Password = hash

	_, err = ctx.repo.CreateUser(u)
	if err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "an unknown exception occurred @todo const", 0)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	session := cache.Session{
		Username: u.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.cache.Set(sessionToken, session); err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "an unknown exception occurred @todo const", 0)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    sessionToken,
		Expires:  time.Now().Add(60 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	models.FormatResponse(w, http.StatusOK, models.DefaultOk(
		models.SessionResponse{
			Username: u.Username,
			Exp:      int(time.Until(session.Expiry).Seconds()),
		},
	),
	)
}

func HashPassword(password string) (string, error) {
	// GenerateFromPassword auto generates salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
