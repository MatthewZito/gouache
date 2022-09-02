package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
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
			// We need to ensure the request body remains readable in the subsequent handler,
			// so we'll initialize a new buffer `b`...
			b := bytes.NewBuffer(make([]byte, 0))

			// TeeReader returns a Reader that writes to `b` what it reads from r.Body.
			reader := io.TeeReader(r.Body, b)

			// Deserialize the reader data into `p`
			json.NewDecoder(reader).Decode(&p)

			// We're done with the original body...
			defer r.Body.Close()

			// NopCloser returns a ReadCloser with a no-op Close method wrapping the provided Reader `r`.
			// We set this to the request and forward it to the subsequent handler.
			r.Body = ioutil.NopCloser(b)
		}

		// Send the report in the background.
		go func() {
			rl := entities.RequestReport{
				Path:       r.RequestURI,
				Method:     method,
				Parameters: p,
			}

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
