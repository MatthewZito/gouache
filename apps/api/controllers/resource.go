package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	cache "github.com/MatthewZito/gouache/cache"
	format "github.com/MatthewZito/gouache/format"
	models "github.com/MatthewZito/gouache/models"
	srv "github.com/MatthewZito/gouache/services"

	"github.com/MatthewZito/turnpike"
)

type ResourceContext struct {
	c *cache.Cache
	l *srv.LoggerClient
}

func NewResourceContext(debug bool) *ResourceContext {
	ctx := &ResourceContext{c: cache.NewCache()}

	if debug {
		ctx.l = srv.NewLogger("resource")
	}

	return ctx
}

func (ctx *ResourceContext) GetResource(w http.ResponseWriter, r *http.Request) {
	key := turnpike.GetParam(r.Context(), "key")
	if key == "" {
		ctx.l.Logf("GetResource - key not provided\n")
		format.FormatError(w, http.StatusBadRequest, "key not provided")
	}

	ctx.l.Logf("GetResource - request key %s\n", key)
	rs := ctx.c.Get(key)

	if v, err := json.Marshal(&rs); err == nil {

		format.FormatResponse(w, http.StatusOK, format.DefaultOk(v))
	} else {
		ctx.l.Logf("GetResource - marshalling error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
	}
}

func (ctx *ResourceContext) GetAllResources(w http.ResponseWriter, r *http.Request) {
	rs := ctx.c.All()

	payload := format.Response{
		Ok:   true,
		Data: rs,
	}

	if v, err := json.Marshal(&payload); err == nil {
		format.FormatResponse(w, http.StatusOK, v)
	} else {
		ctx.l.Logf("GetAllResources - marshalling error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
	}
}

func (ctx *ResourceContext) AddResource(w http.ResponseWriter, r *http.Request) {
	rs := models.Resource{}

	if err := json.NewDecoder(r.Body).Decode(&rs); err != nil {
		ctx.l.Logf("AddResource - decoding error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx.l.Logf("AddResource - put key: %s value: %v expires: %d\n", rs.Key, rs.Value, rs.Expires)

	ctx.c.Put(rs.Key, rs.Value, rs.Expires)
	format.FormatResponse(w, http.StatusOK, format.DefaultOk(nil))
}

func (ctx *ResourceContext) UpdateResource(w http.ResponseWriter, r *http.Request) {
	key := turnpike.GetParam(r.Context(), "key")
	if key == "" {
		ctx.l.Logf("UpdateResource - key not provided\n")
		format.FormatError(w, http.StatusBadRequest, "key not provided")
	}

	rs := models.Resource{}

	if err := json.NewDecoder(r.Body).Decode(&rs); err != nil {
		ctx.l.Logf("UpdateResource - decoding error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx.l.Logf("UpdateResource - put key: %s value: %v expires: %d\n", rs.Key, rs.Value, rs.Expires)
	ctx.c.Put(key, rs.Value, rs.Expires)
	format.FormatResponse(w, http.StatusOK, format.DefaultOk(nil))
}

func (ctx *ResourceContext) DeleteResource(w http.ResponseWriter, r *http.Request) {
	key := turnpike.GetParam(r.Context(), "key")
	ctx.l.Logf("DeleteResource - request key %s\n", key)

	if ok := ctx.c.Delete(key); !ok {
		ctx.l.Logf("DeleteResource - delete key %s failed\n", key)
		format.FormatError(w, http.StatusBadRequest, fmt.Sprintf("delete failed for key %s", key))
	}

	format.FormatResponse(w, http.StatusOK, format.DefaultOk(nil))
}
