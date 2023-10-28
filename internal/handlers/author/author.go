package author

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/doublehops/dhapi-example/internal/handlers"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository/repositoryauthor"
	req "github.com/doublehops/dhapi-example/internal/request"
	"github.com/doublehops/dhapi-example/internal/service"
)

type Handle struct {
	repo *repositoryauthor.Author
	srv  *service.AuthorService
	base *handlers.BaseHandler
}

func New(app *service.App) *Handle {
	ar := repositoryauthor.New(app.Log)

	return &Handle{
		repo: ar,
		srv:  service.New(app, ar),
		base: &handlers.BaseHandler{
			Log: app.Log,
		},
	}
}

func (h *Handle) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.srv.Log.Info(c, "Request made to CreateAuthor", nil)

	author := &model.Author{}
	if err := json.NewDecoder(r.Body).Decode(author); err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, req.UnableToParseResp())

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := req.GetValidateErrResp(errors, req.ValidationError.Error())
		h.base.WriteJson(c, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.srv.Create(c, author)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, req.GeneralErrResp(req.ErrorProcessingRequest.Error()))

		return
	}

	h.base.WriteJson(c, w, http.StatusOK, req.GetSingleItemResp(a))
}

func (h *Handle) UpdateByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.base.GetUser(c)
	h.srv.Log.Info(c, "Request made to UpdateAuthor", nil)

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.srv.GetByID(c, author, int32(i))
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, err)

		return
	}

	if author.ID == 0 {
		h.base.WriteJson(c, w, http.StatusNotFound, req.GetNotFoundResp())

		return
	}

	if !h.srv.HasPermission(userID, author) {
		h.base.WriteJson(c, w, http.StatusForbidden, req.GetNotAuthorisedResp())

		return
	}

	if err := json.NewDecoder(r.Body).Decode(author); err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, req.UnableToParseResp())

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := req.GetValidateErrResp(errors, req.ValidationError.Error())
		h.base.WriteJson(c, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.srv.Update(c, author)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.base.WriteJson(c, w, http.StatusOK, req.GetSingleItemResp(a))
}

func (h *Handle) DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.base.GetUser(c)
	h.srv.Log.Info(c, "Request made to DELETE author", nil)

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.srv.GetByID(c, author, int32(i))
	if err != nil {
		h.base.WriteJson(c, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if author.ID == 0 {
		h.base.WriteJson(c, w, http.StatusNotFound, req.GetNotFoundResp())

		return
	}

	if !h.srv.HasPermission(userID, author) {
		h.base.WriteJson(c, w, http.StatusForbidden, req.GetNotAuthorisedResp())

		return
	}

	if err = h.srv.DeleteByID(c, author, int32(i)); err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, req.ErrorProcessingRequestResp())

		return
	}

	h.base.WriteJson(c, w, http.StatusNoContent, nil)
}

func (h *Handle) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	userID := h.base.GetUser(c)
	h.srv.Log.Info(c, "Request made to Get author", nil)

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	err = h.srv.GetByID(c, author, int32(i))
	if err != nil {
		h.base.WriteJson(c, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if author.ID == 0 {
		h.base.WriteJson(c, w, http.StatusNotFound, req.GetNotFoundResp())

		return
	}

	if !h.srv.HasPermission(userID, author) {
		h.base.WriteJson(c, w, http.StatusForbidden, req.GetNotAuthorisedResp())

		return
	}

	h.base.WriteJson(c, w, http.StatusOK, req.GetSingleItemResp(author))
}

func (h *Handle) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := r.Context()
	h.srv.Log.Info(c, "Request made to Get authors", nil)

	p := req.GetPaginationReq(r)

	authors, err := h.srv.GetAll(c, p)
	if err != nil {
		h.base.WriteJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.base.WriteJson(c, w, http.StatusOK, req.GetListResp(authors, p))
}
