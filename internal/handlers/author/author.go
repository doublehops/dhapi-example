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
	app *app.App
	ar  *repositoryauthor.RepositoryAuthor
	as  *service.AuthorService
}

func New(app *app.App) *Handle {
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

	var author *model.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "Unable to parse request")

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := resp.GetValidateErrResp(errors, resp.ValidationError.Error())
		h.writeJson(c, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.as.Create(c, author)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
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
	h.app.Log.Info(c, "Request made to UpdateAuthor")

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author := &model.Author{}
	if err := json.NewDecoder(r.Body).Decode(author); err != nil {
		errResp := resp.GetValidateErrResp(nil, "Unable to parse request")
		h.writeJson(c, w, http.StatusBadRequest, errResp)

		return
	}

	if errors := author.Validate(); len(errors) > 0 {
		errs := resp.GetValidateErrResp(errors, resp.ValidationError.Error())
		h.writeJson(c, w, http.StatusBadRequest, errs)

		return
	}

	err = h.as.GetByID(c, author, int32(i))
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, err)

		return
	}

	if author.ID == 0 {
		h.writeJson(c, w, http.StatusNotFound, resp.GetNotFoundResp())

		return
	}

	// todo - check authorised.

	a, err := h.as.Update(c, author)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
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

	// todo - check authorised.

	h.as.DeleteByID(c, author, int32(i))

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(author))
}

func (h *Handle) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
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

	// todo - check authorised.

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
