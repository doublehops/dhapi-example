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
	ar := repositoryauthor.New(app.DB, app.Log)
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

	a, err := h.as.Create(c, author)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
}

// todo - move this to a reusable place.
func (h *Handle) writeJson(ctx context.Context, w http.ResponseWriter, statusCode int, res interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		h.app.Log.Error(ctx, "unable to marshal to JSON. T%s"+err.Error())
	}
}

func (h *Handle) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.app.Log.Info(c, "Request made to UpdateAuthor")

	ID := ps.ByName("id")

	var author *model.Author
	err := json.NewEncoder(w).Encode(author)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "Unable to parse request")

		return
	}

	i, err := strconv.Atoi(ID)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author.ID = int32(i)

	a, err := h.as.Update(c, author)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.writeJson(c, w, http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.app.Log.Info(c, "Request made to Get author")

	ID := ps.ByName("id")

	author := &model.Author{}

	i, err := strconv.Atoi(ID)
	if err != nil {
		h.writeJson(c, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	intID := int32(i)

	err = h.as.GetByID(c, intID, author)
	if err != nil {
		h.writeJson(c, w, http.StatusInternalServerError, "Unable to process request")

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
