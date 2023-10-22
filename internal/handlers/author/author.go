package author

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/doublehops/dhapi/resp"

	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository/repositoryauthor"
	"github.com/doublehops/dhapi-example/internal/service"
)

type Handle struct {
	app *service.App
	ar  *repositoryauthor.RepositoryAuthor
	as  *service.AuthorService
}

func New(app *service.App) *Handle {
	ar := repositoryauthor.New(app.Log)

	return &Handle{
		app: app,
		ar:  ar,
		as:  service.New(app, ar),
	}
}

func (h *Handle) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.app.Log.Info(c, "Request made to CreateAuthor")

	author := &model.Author{}
	if err := json.NewDecoder(r.Body).Decode(author); err != nil {
		h.writeJson(c, w, http.StatusBadRequest, resp.UnableToParseResp())

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := resp.GetValidateErrResp(errors, resp.ValidationError.Error())
		h.writeJson(c, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.as.Create(c, author)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, resp.GeneralErrResp(resp.ErrorProcessingRequest.Error()))

		return
	}

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) GetUser(ctx context.Context) int32 {
	var intValue int32
	var ok bool

	val := ctx.Value(app.UserIDKey)
	if intValue, ok = val.(int32); ok {

	}

	return intValue
}

// todo - move this to a reusable place.
func (h *Handle) writeJson(ctx context.Context, w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		h.app.Log.Error(ctx, "unable to marshal to JSON. "+err.Error())
	}
}

func (h *Handle) UpdateByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.GetUser(c)
	h.app.Log.Info(c, "Request made to UpdateAuthor")

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.as.GetByID(c, author, int32(i))
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, err)

		return
	}

	if author.ID == 0 {
		h.writeJson(c, w, http.StatusNotFound, resp.GetNotFoundResp())

		return
	}

	if !h.app.HasPermission(userID, author) {
		h.writeJson(c, w, http.StatusForbidden, resp.GetNotAuthorisedResp())

		return
	}

	if err := json.NewDecoder(r.Body).Decode(author); err != nil {
		h.writeJson(c, w, http.StatusBadRequest, resp.UnableToParseResp())

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := resp.GetValidateErrResp(errors, resp.ValidationError.Error())
		h.writeJson(c, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.as.Update(c, author)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.GetUser(c)
	h.app.Log.Info(c, "Request made to DELETE author")

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.as.GetByID(c, author, int32(i))
	if err != nil {
		h.writeJson(c, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if author.ID == 0 {
		h.writeJson(c, w, http.StatusNotFound, resp.GetNotFoundResp())

		return
	}

	if !h.app.HasPermission(userID, author) {
		h.writeJson(c, w, http.StatusForbidden, resp.GetNotAuthorisedResp())

		return
	}

	if err = h.as.DeleteByID(c, author, int32(i)); err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, resp.ErrorProcessingRequestResp())

		return
	}

	h.writeJson(c, w, http.StatusNoContent, nil)
}

func (h *Handle) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.GetUser(c)
	h.app.Log.Info(c, "Request made to Get author")

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.as.GetByID(c, author, int32(i))
	if err != nil {
		h.writeJson(c, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if author.ID == 0 {
		h.writeJson(c, w, http.StatusNotFound, resp.GetNotFoundResp())

		return
	}

	if !h.app.HasPermission(userID, author) {
		h.writeJson(c, w, http.StatusForbidden, resp.GetNotAuthorisedResp())

		return
	}

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(author))
}

func (h *Handle) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := r.Context()
	h.app.Log.Info(c, "Request made to Get authors")

	authors, err := h.as.GetAll(c)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	pagination := resp.Pagination{}

	h.writeJson(c, w, http.StatusOK, resp.GetListResp(authors, pagination))
}
