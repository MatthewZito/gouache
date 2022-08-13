package controllers

import (
	"fmt"
	"net/http"
	"time"

	format "github.com/exbotanical/gouache/format"
	srv "github.com/exbotanical/gouache/services"
)

// MetaContext holds shared context for meta endpoints.
type MetaContext struct {
	l *srv.LoggerClient
}

// NewMetaContext creates a new MetaContext object.
func NewMetaContext(debug bool) *MetaContext {
	ctx := &MetaContext{}

	if debug {
		ctx.l = srv.NewLogger("meta")
	}

	return ctx
}

// GetTime
func (ctx *MetaContext) GetTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixNano()
	ctx.l.Logf("GetTime %d", now)
	format.FormatResponse(w, http.StatusOK, format.DefaultOk(fmt.Sprint(now)))
}
