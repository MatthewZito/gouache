package controllers

import (
	"encoding/json"
	"net/http"

	db "github.com/MatthewZito/gouache/db"
	format "github.com/MatthewZito/gouache/format"
	models "github.com/MatthewZito/gouache/models"
	srv "github.com/MatthewZito/gouache/services"

	"github.com/MatthewZito/turnpike"
)

type ResourceContext struct {
	db *db.DB
	l  *srv.LoggerClient
}

func NewResourceContext(debug bool, db *db.DB) *ResourceContext {
	ctx := &ResourceContext{db: db}

	if debug {
		ctx.l = srv.NewLogger("resource")
	}

	return ctx
}

func (ctx *ResourceContext) GetResource(w http.ResponseWriter, r *http.Request) {
	id := turnpike.GetParam(r.Context(), "id")
	if id == "" {
		ctx.l.Logf("GetResource - id not provided\n")
		format.FormatError(w, http.StatusBadRequest, "id not provided")
	}

	ctx.l.Logf("GetResource - request key %s\n", id)

	if r, err := ctx.db.GetResource(id); err == nil {
		format.FormatResponse(w, http.StatusOK, format.DefaultOk(r))
	} else {
		ctx.l.Logf("GetResource - database error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
	}
}

func (ctx *ResourceContext) GetAllResources(w http.ResponseWriter, r *http.Request) {
	if rs, err := ctx.db.GetResources(); err == nil {
		if v, err := json.Marshal(&format.Response{
			Ok:   true,
			Data: rs,
		}); err == nil {
			format.FormatResponse(w, http.StatusOK, v)
		} else {
			ctx.l.Logf("GetAllResources - marshalling error %v\n", err)
			format.FormatError(w, http.StatusBadRequest, err.Error())
		}
	} else {
		ctx.l.Logf("GetAllResources - database error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
	}
}

func (ctx *ResourceContext) CreateResource(w http.ResponseWriter, r *http.Request) {
	rs := &models.NewResourceTemplate{}

	if err := json.NewDecoder(r.Body).Decode(&rs); err != nil {
		ctx.l.Logf("CreateResource - decoding error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	if id, err := ctx.db.CreateResource(rs); err != nil {
		ctx.l.Logf("CreateResource - database error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
	} else {
		ctx.l.Logf("CreateResource - %v with id %s\n", rs, id)
		format.FormatResponse(w, http.StatusOK, format.DefaultOk(id))
	}
}

func (ctx *ResourceContext) UpdateResource(w http.ResponseWriter, r *http.Request) {
	id := turnpike.GetParam(r.Context(), "id")
	if id == "" {
		ctx.l.Logf("UpdateResource - id not provided\n")
		format.FormatError(w, http.StatusBadRequest, "id not provided")
	}

	rs := &models.UpdateResourceTemplate{}

	if err := json.NewDecoder(r.Body).Decode(&rs); err != nil {
		ctx.l.Logf("UpdateResource - decoding error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
		return
	}

	rs.Id = id
	if err := ctx.db.UpdateResource(rs); err != nil {
		ctx.l.Logf("UpdateResource - database error %v\n", err)
		format.FormatError(w, http.StatusBadRequest, err.Error())
	} else {
		ctx.l.Logf("UpdateResource - %v\n", rs)
		format.FormatResponse(w, http.StatusOK, format.DefaultOk(nil))
	}
}

// func (ctx *ResourceContext) DeleteResource(w http.ResponseWriter, r *http.Request) {
// 	key := turnpike.GetParam(r.Context(), "key")
// 	ctx.l.Logf("DeleteResource - request key %s\n", key)

// 	if ok := ctx.c.Delete(key); !ok {
// 		ctx.l.Logf("DeleteResource - delete key %s failed\n", key)
// 		format.FormatError(w, http.StatusBadRequest, fmt.Sprintf("delete failed for key %s", key))
// 	}

// 	format.FormatResponse(w, http.StatusOK, format.DefaultOk(nil))
// }
