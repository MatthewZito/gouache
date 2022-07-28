package controllers

import (
	"fmt"
	"net/http"
	"time"

	format "github.com/MatthewZito/gouache/format"
	srv "github.com/MatthewZito/gouache/services"
)

type MetaContext struct {
	l *srv.LoggerClient
}

func NewMetaContext(debug bool) *MetaContext {
	ctx := &MetaContext{}

	if debug {
		ctx.l = srv.NewLogger("meta")
	}

	return ctx
}

func (ctx *MetaContext) GetTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixNano()
	ctx.l.Logf("GetTime %d", now)
	format.FormatResponse(w, http.StatusOK, format.DefaultOk(fmt.Sprint(now)))
}
