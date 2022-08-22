package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/exbotanical/gouache/cache"
	"github.com/exbotanical/gouache/db"
	"github.com/exbotanical/gouache/format"
	"github.com/exbotanical/gouache/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const cookieId = "gouache_session"

type SessionContext struct {
	cache cache.SerializableStore
	db    *db.DB
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewSessionContext(client cache.SerializableStore, db *db.DB) SessionContext {
	return SessionContext{cache: client, db: db}
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
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	u, err := ctx.db.GetUser(credentials.Username)
	if err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !CheckPasswordHash(credentials.Password, u.Password) {
		format.FormatError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	session := cache.Session{
		Username: credentials.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.cache.Set(sessionToken, session); err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	format.FormatResponse(w, http.StatusOK, format.DefaultOk(
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
		format.FormatError(w, http.StatusUnauthorized, err.Error())
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

	format.FormatResponse(w, http.StatusOK, format.DefaultOk(nil))
}

func (ctx SessionContext) RenewSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieId)
	if err != nil {
		format.FormatError(w, http.StatusUnauthorized, err.Error())
		return
	}

	sessionToken := cookie.Value

	session, err := ctx.cache.Get(sessionToken)
	if err != nil {
		format.FormatError(w, http.StatusUnauthorized, err.Error())
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

	format.FormatResponse(w, http.StatusOK, format.DefaultOk(models.SessionResponse{Username: session.Username, Exp: int(time.Until(session.Expiry).Seconds())}))
}

// @todo Determine when user exists
func (ctx SessionContext) Register(w http.ResponseWriter, r *http.Request) {
	u := &models.NewUserTemplate{}
	fmt.Println("REGISTER")

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := HashPassword(u.Password)
	if err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.Password = hash

	_, err = ctx.db.CreateUser(u)
	if err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	session := cache.Session{
		Username: u.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.cache.Set(sessionToken, session); err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieId,
		Value:    sessionToken,
		Expires:  time.Now().Add(60 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	format.FormatResponse(w, http.StatusOK, format.DefaultOk(
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
