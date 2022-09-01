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
func (ctx AuthProvider) Authorize(next http.Handler) http.Handler {
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
			session, err := ctx.ss.GetSession(token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else if session.IsExpired() {
				ctx.ss.DeleteSession(token)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// ReportRequest is a middleware that sends reports of all inbound requests to the gouache/reporting service.
func (ctx AuthProvider) ReportRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		if utils.Contains([]string{http.MethodOptions, http.MethodHead, http.MethodTrace, http.MethodConnect}, method) {

			next.ServeHTTP(w, r)
			return
		}

		var p map[string]interface{}

		if method != http.MethodGet {
			// data, _ := ioutil.ReadAll(r.Body)
			// @todo restore
			json.NewDecoder(r.Body).Decode(&p)

			// r.Body.Close()
			// r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		rl := entities.RequestReport{
			Path:       r.RequestURI,
			Method:     method,
			Parameters: p,
		}

		go func() {
			ctx.rs.SendReport(context.TODO(), entities.HTTP_REQUEST_RECV, rl)
		}()

		next.ServeHTTP(w, r)
	})
}

// handleException handles http controller exceptions and reports the error to gouache/reporting.
func (ctx AuthProvider) handleException(w http.ResponseWriter, r *http.Request, status int, internal string, friendly string) {
	models.SendGouacheException(w, status, internal, friendly, 0)
	ctx.rs.SendControllerErrorReport(r, internal, friendly)
}
