package author

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/doublehops/dhapi/resp"

	"github.com/doublehops/dhapi-example/internal/handlers"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository/repositoryauthor"
	"github.com/doublehops/dhapi-example/internal/service"
)

type Handle struct {
	ar   *repositoryauthor.RepositoryAuthor
	as   *service.AuthorService
	base *handlers.BaseHandler
}

func New(app *service.App) *Handle {
	ar := repositoryauthor.New(app.Log)

	return &Handle{
		ar: ar,
		as: service.New(app, ar),
		base: &handlers.BaseHandler{
			Log: app.Log,
		},
	}
}

func (h *Handle) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.as.Log.Info(c, "Request made to CreateAuthor")

	author := &model.Author{}
	if err := json.NewDecoder(r.Body).Decode(author); err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, resp.UnableToParseResp())

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := resp.GetValidateErrResp(errors, resp.ValidationError.Error())
		h.base.WriteJson(c, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.as.Create(c, author)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, resp.GeneralErrResp(resp.ErrorProcessingRequest.Error()))

		return
	}

	h.base.WriteJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) UpdateByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.base.GetUser(c)
	h.as.Log.Info(c, "Request made to UpdateAuthor")

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.as.GetByID(c, author, int32(i))
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, err)

		return
	}

	if author.ID == 0 {
		h.base.WriteJson(c, w, http.StatusNotFound, resp.GetNotFoundResp())

		return
	}

	if !h.as.HasPermission(userID, author) {
		h.base.WriteJson(c, w, http.StatusForbidden, resp.GetNotAuthorisedResp())

		return
	}

	if err := json.NewDecoder(r.Body).Decode(author); err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, resp.UnableToParseResp())

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := resp.GetValidateErrResp(errors, resp.ValidationError.Error())
		h.base.WriteJson(c, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.as.Update(c, author)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.base.WriteJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.base.GetUser(c)
	h.as.Log.Info(c, "Request made to DELETE author")

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.as.GetByID(c, author, int32(i))
	if err != nil {
		h.base.WriteJson(c, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if author.ID == 0 {
		h.base.WriteJson(c, w, http.StatusNotFound, resp.GetNotFoundResp())

		return
	}

	if !h.as.HasPermission(userID, author) {
		h.base.WriteJson(c, w, http.StatusForbidden, resp.GetNotAuthorisedResp())

		return
	}

	if err = h.as.DeleteByID(c, author, int32(i)); err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, resp.ErrorProcessingRequestResp())

		return
	}

	h.base.WriteJson(c, w, http.StatusNoContent, nil)
}

func (h *Handle) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.base.GetUser(c)
	h.as.Log.Info(c, "Request made to Get author")

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.as.GetByID(c, author, int32(i))
	if err != nil {
		h.base.WriteJson(c, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if author.ID == 0 {
		h.base.WriteJson(c, w, http.StatusNotFound, resp.GetNotFoundResp())

		return
	}

	if !h.as.HasPermission(userID, author) {
		h.base.WriteJson(c, w, http.StatusForbidden, resp.GetNotAuthorisedResp())

		return
	}

	h.base.WriteJson(c, w, http.StatusOK, resp.GetSingleItemResp(author))
}

func (h *Handle) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := r.Context()
	h.as.Log.Info(c, "Request made to Get authors")

	authors, err := h.as.GetAll(c)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	pagination := resp.Pagination{}

	h.base.WriteJson(c, w, http.StatusOK, resp.GetListResp(authors, pagination))
}
