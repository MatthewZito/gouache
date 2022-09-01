package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	entities "github.com/exbotanical/gouache/entities/reporting"
	"github.com/exbotanical/gouache/models"
	"github.com/exbotanical/gouache/utils"
)

// Authorize is a middleware that checks the user's session cookie to evaluate whether they're authorized to access the system.
func (ctx SessionContext) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(COOKIE_ID)

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

// ReportRequest is a middleware that sends reports of all inbound requests to the gouache/reporting service.
func (ctx SessionContext) ReportRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		if utils.Contains([]string{http.MethodOptions, http.MethodHead, http.MethodTrace, http.MethodConnect}, method) {

			next.ServeHTTP(w, r)
			return
		}

		var p map[string]interface{}

		if method != http.MethodGet {
			json.NewDecoder(r.Body).Decode(&p)
		}

		rl := entities.RequestReport{
			Path:       r.RequestURI,
			Method:     method,
			Parameters: p,
		}

		ctx.q.SendReport(context.TODO(), entities.HTTP_REQUEST_RECV, rl)

		next.ServeHTTP(w, r)
	})
}

// handleException handles http controller exceptions and reports the error to gouache/reporting.
func (ctx SessionContext) handleException(w http.ResponseWriter, r *http.Request, status int, internal string, friendly string) {
	models.SendGouacheException(w, status, internal, friendly, 0)
	ctx.q.SendControllerErrorReport(r, internal, friendly)
}
