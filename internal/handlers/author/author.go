package author

import (
	"encoding/json"
	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository/repositoryauthor"
	"github.com/doublehops/dhapi-example/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"

	"github.com/doublehops/dhapi/resp"
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
	err := json.Unmarshal(r.Body, &author)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable to parse request")

		return
	}

	a, err := h.as.Create(c, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to process request")

		return
	}

	c.JSON(http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.app.Log.Info(c, "Request made to UpdateAuthor")

	ID := c.Param("id")

	var author *model.Author
	//fmt.Printf("body: %s", c.Request.Body)
	err := c.BindJSON(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable to parse request")

		return
	}

	i, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID is not a valid value")

		return
	}

	author.ID = int32(i)

	a, err := h.as.Update(c, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to process request")

		return
	}

	c.JSON(http.StatusOK, resp.GetSingleItemResp(a))
}

func (h *Handle) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.app.Log.Info(c, "Request made to Get author")

	ID := c.Param("id")

	author := &model.Author{}

	i, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID is not a valid value")

		return
	}

	intID := int32(i)

	err = h.as.GetByID(c, intID, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to process request")

		return
	}

	c.JSON(http.StatusOK, resp.GetSingleItemResp(author))
}

func (h *Handle) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := r.Context()
	h.app.Log.Info(c, "Request made to Get authors")

	authors, err := h.as.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to process request")

		return
	}

	pagination := resp.Pagination{}

	c.JSON(http.StatusOK, resp.GetListResp(authors, pagination))
}
