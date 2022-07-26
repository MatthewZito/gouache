package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	cache "github.com/MatthewZito/gouache/cache"
	format "github.com/MatthewZito/gouache/format"
	"github.com/MatthewZito/gouache/premux"
)

type Resource struct {
	Key     string
	Value   interface{}
	Expires int64
}

type ResourceCache struct {
	c *cache.Cache
}

func NewResourceCache() *ResourceCache {
	return &ResourceCache{c: cache.NewCache()}
}

func (rc *ResourceCache) GetResource(w http.ResponseWriter, r *http.Request) {
	key := premux.GetParam(r.Context(), "key")
	if key == "" {
		format.FormatError(w, http.StatusBadRequest, "missing key")
	}

	rs := rc.c.Get(key)
	if v, err := json.Marshal(&rs); err == nil {
		format.FormatResponse(w, http.StatusOK, map[string]string{"ok": "true", "rs": string(v)})
	} else {
		format.FormatError(w, http.StatusBadRequest, err.Error())
	}
}

func (rc *ResourceCache) AddResource(w http.ResponseWriter, r *http.Request) {
	rs := Resource{}

	err := json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	rc.c.Put(rs.Key, rs.Value, rs.Expires)
	format.FormatResponse(w, http.StatusOK, format.DefaultOk())
}

func (rc *ResourceCache) UpdateResource(w http.ResponseWriter, r *http.Request) {
	key := premux.GetParam(r.Context(), "key")
	if key == "" {
		format.FormatError(w, http.StatusBadRequest, "missing key")
	}

	rs := Resource{}

	err := json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	rc.c.Put(key, rs.Value, rs.Expires)
	format.FormatResponse(w, http.StatusOK, format.DefaultOk())
}

func (rc *ResourceCache) DeleteResource(w http.ResponseWriter, r *http.Request) {
	key := premux.GetParam(r.Context(), "key")

	if ok := rc.c.Delete(key); !ok {
		format.FormatError(w, http.StatusBadRequest, fmt.Sprintf("delete failed for key %s", key))
	}

	format.FormatResponse(w, http.StatusOK, format.DefaultOk())
}
