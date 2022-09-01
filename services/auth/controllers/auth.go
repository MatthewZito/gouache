package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/exbotanical/gouache/entities"
	"github.com/exbotanical/gouache/models"
	"github.com/exbotanical/gouache/services"
	"github.com/exbotanical/gouache/utils"

	"github.com/google/uuid"
)

// COOKIE_ID represents the session identifier.
const COOKIE_ID = "gouache_session"

// AuthProvider holds the auth controller context, including the redis session cache and dynamodb client.
type AuthProvider struct {
	ss services.SessionService
	us services.UserService
	rs services.ReportService
}

// Credentials represents a user's input credentials.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewAuthProvider initializes a new AuthProvider object.
func NewAuthProvider(ss services.SessionService, us services.UserService, rs services.ReportService) AuthProvider {
	return AuthProvider{ss: ss, us: us, rs: rs}
}

// Login authenticates a user given correct `Credentials`.
func (ctx AuthProvider) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("INVOKE")
	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		fmt.Println(err)
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "invalid credentials")
		return
	}

	u, err := ctx.us.GetUser(credentials.Username)
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

	session := entities.Session{
		Username: credentials.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.ss.SetSession(sessionToken, session); err != nil {
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
func (ctx AuthProvider) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(COOKIE_ID)
	if err != nil {
		ctx.handleException(w, r, http.StatusUnauthorized, err.Error(), "an unknown exception occurred @todo const")
		return
	}

	sessionToken := cookie.Value

	ctx.ss.DeleteSession(sessionToken)

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
func (ctx AuthProvider) RenewSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(COOKIE_ID)
	if err != nil {
		ctx.handleException(w, r, http.StatusUnauthorized, err.Error(), "unauthorized")
		return
	}

	sessionToken := cookie.Value

	session, err := ctx.ss.GetSession(sessionToken)
	if err != nil {
		ctx.handleException(w, r, http.StatusUnauthorized, err.Error(), "unauthorized")
		return
	}

	expiresAt := time.Now().Add(60 * time.Minute)
	session.Expiry = expiresAt

	ctx.ss.SetSession(sessionToken, *session)

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
func (ctx AuthProvider) Register(w http.ResponseWriter, r *http.Request) {
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

	err = ctx.us.CreateUser(u)
	if err != nil {
		ctx.handleException(w, r, http.StatusBadRequest, err.Error(), "an unknown exception occurred @todo const")
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	session := entities.Session{
		Username: u.Username,
		Expiry:   expiresAt,
	}

	if err := ctx.ss.SetSession(sessionToken, session); err != nil {
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
