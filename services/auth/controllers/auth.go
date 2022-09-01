package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/exbotanical/gouache/entities"
	"github.com/exbotanical/gouache/models"
	"github.com/exbotanical/gouache/repositories"
	"github.com/exbotanical/gouache/utils"

	"github.com/google/uuid"
)

// COOKIE_ID represents the session identifier.
const COOKIE_ID = "gouache_session"

// SessionContext holds the auth controller context, including the redis session cache and dynamodb client.
type SessionContext struct {
	cache repositories.SessionManager
	repo  *repositories.UserTable
	q     *repositories.ReportRepository
}

// Credentials represents a user's input credentials.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewSessionContext initializes a new SessionContext object.
func NewSessionContext(client repositories.SessionManager, repo *repositories.UserTable, q *repositories.ReportRepository) SessionContext {
	return SessionContext{cache: client, repo: repo, q: q}
}

// Login authenticates a user given correct `Credentials`.
func (ctx SessionContext) Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "invalid credentials")
		return
	}

	u, err := ctx.repo.GetUser(credentials.Username)
	if err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "invalid credentials")
		return
	}

	if !utils.CheckPasswordHash(credentials.Password, u.Password) {
		ctx.handleException(w, r, http.StatusUnauthorized, "password mismatch", "invalid credentials")
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	session := repositories.Session{
		Username: credentials.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.cache.Set(sessionToken, session); err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "an unknown exception occurred @todo const")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     COOKIE_ID,
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	models.SendGouacheResponse(w, http.StatusOK, models.ToOk(
		entities.SessionResponse{
			Username: u.Username,
			Exp:      int(time.Until(session.Expiry).Seconds()),
		},
	),
	)
}

// Logout destroys the user's session and removes their session cookie.
func (ctx SessionContext) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(COOKIE_ID)
	if err != nil {
		ctx.handleException(w, r, http.StatusUnauthorized, err.Error(), "an unknown exception occurred @todo const")
		return
	}

	sessionToken := cookie.Value

	ctx.cache.Delete(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:     COOKIE_ID,
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})

	models.SendGouacheResponse(w, http.StatusOK, models.ToOk(nil))
}

// RenewSession renews the user's session cookie.
func (ctx SessionContext) RenewSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(COOKIE_ID)
	if err != nil {
		ctx.handleException(w, r, http.StatusUnauthorized, err.Error(), "unauthorized")
		return
	}

	sessionToken := cookie.Value

	session, err := ctx.cache.Get(sessionToken)
	if err != nil {
		ctx.handleException(w, r, http.StatusUnauthorized, err.Error(), "unauthorized")
		return
	}

	expiresAt := time.Now().Add(60 * time.Minute)
	session.Expiry = expiresAt

	ctx.cache.Set(sessionToken, *session)

	http.SetCookie(w, &http.Cookie{
		Name:     COOKIE_ID,
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	models.SendGouacheResponse(
		w,
		http.StatusOK,
		models.ToOk(entities.SessionResponse{
			Username: session.Username,
			Exp:      int(time.Until(session.Expiry).Seconds()),
		}))
}

// Register creates a new user and corresponding session when provided a `NewUserModel`.
// @todo Determine when user exists
func (ctx SessionContext) Register(w http.ResponseWriter, r *http.Request) {
	u := models.NewUserModel{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "invalid credentials provided")
		return
	}

	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "invalid credentials provided")
		return
	}

	u.Password = hash

	err = ctx.repo.CreateUser(u)
	if err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "an unknown exception occurred @todo const")
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	session := repositories.Session{
		Username: u.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.cache.Set(sessionToken, session); err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "an unknown exception occurred @todo const")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     COOKIE_ID,
		Value:    sessionToken,
		Expires:  time.Now().Add(60 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	models.SendGouacheResponse(w, http.StatusOK, models.ToOk(
		entities.SessionResponse{
			Username: u.Username,
			Exp:      int(time.Until(session.Expiry).Seconds()),
		},
	),
	)
}
