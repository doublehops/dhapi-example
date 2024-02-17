package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/doublehops/dhapi-example/internal/model/mynewtable"
	"github.com/doublehops/dhapi-example/internal/repository/mynewtable"
	req "github.com/doublehops/dhapi-example/internal/request"
	"github.com/doublehops/dhapi-example/internal/service"
	"github.com/doublehops/dhapi-example/internal/tools"
)

type Handle struct {
	repo *repository.MyNewTable
	srv  *service.MyNewTableService
	base *BaseHandler
}

func New(app *service.App) *Handle {
	repo := mynewtable.New(app.Log)

	return &Handle{
		repo: repo,
		srv:  service.New(app, repo),
		base: &BaseHandler{
			Log: app.Log,
		},
	}
}
12
func (h *Handle) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	h.base.Log.Info(ctx, "Request made to "+tools.CurrentFunction(), nil)

	record := &model.MyNewTable{}
	if err := json.NewDecoder(r.Body).Decode(record); err != nil {
		h.base.WriteJson(ctx, w, http.StatusBadRequest, req.UnableToParseResp())

		return
	}

	if errors := record.Validate(); len(errors) > 0 {
		errs := req.GetValidateErrResp(errors, req.ValidationError.Error())
		h.base.WriteJson(ctx, w, http.StatusBadRequest, errs)

		return
	}

	a, err := h.srv.Create(ctx, record)
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusInternalServerError, req.GeneralErrResp(req.ErrorProcessingRequest.Error()))

		return
	}

	h.base.WriteJson(ctx, w, http.StatusOK, req.GetSingleItemResp(a))
}

func (h *Handle) UpdateByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	userID := h.base.GetUser(ctx)
	h.base.Log.Info(ctx, "Request made to UpdateMyNewTable", nil)

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	record := &model.MyNewTable{}
	err = h.srv.GetByID(ctx, record, int32(i))
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusBadRequest, err)

		return
	}

	if record.ID == 0 {
		h.base.WriteJson(ctx, w, http.StatusNotFound, req.GetNotFoundResp())

		return
	}

	if !h.srv.HasPermission(userID, record) {
		h.base.WriteJson(ctx, w, http.StatusForbidden, req.GetNotAuthorisedResp())

		return
	}

	if err := json.NewDecoder(r.Body).Decode(record); err != nil {
		h.base.WriteJson(ctx, w, http.StatusBadRequest, req.UnableToParseResp())

		return
	}

	if errors := record.Validate(); len(errors) > 0 {
		errs := req.GetValidateErrResp(errors, req.ValidationError.Error())
		h.base.WriteJson(ctx, w, http.StatusBadRequest, errs)

		return
	}

	err := h.srv.Update(ctx, record)
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.base.WriteJson(ctx, w, http.StatusOK, req.GetSingleItemResp(record))
}

func (h *Handle) DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	userID := h.base.GetUser(ctx)
	h.base.Log.Info(ctx, "Request made to DELETE myNewTable", nil)

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	record := &model.MyNewTable{}
	err = h.srv.GetByID(ctx, record, int32(i))
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if record.ID == 0 {
		h.base.WriteJson(ctx, w, http.StatusNotFound, req.GetNotFoundResp())

		return
	}

	if !h.srv.HasPermission(userID, record) {
		h.base.WriteJson(ctx, w, http.StatusForbidden, req.GetNotAuthorisedResp())

		return
	}

	if err = h.srv.DeleteByID(ctx, record, int32(i)); err != nil {
		h.base.WriteJson(ctx, w, http.StatusInternalServerError, req.ErrorProcessingRequestResp())

		return
	}

	h.base.WriteJson(ctx, w, http.StatusNoContent, nil)
}

func (h *Handle) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	userID := h.base.GetUser(ctx)
	h.base.Log.Info(ctx, "Request made to Get myNewTable", nil)

	ID := ps.ByName("id")
	i, err := strconv.Atoi(ID)
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusBadRequest, "ID is not a valid value")

		return
	}

	record := &model.MyNewTable{}
	err = h.srv.GetByID(ctx, record, int32(i))
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusNotFound, "Unable to find record")

		return
	}

	if record.ID == 0 {
		h.base.WriteJson(ctx, w, http.StatusNotFound, req.GetNotFoundResp())

		return
	}

	if !h.srv.HasPermission(userID, record) {
		h.base.WriteJson(ctx, w, http.StatusForbidden, req.GetNotAuthorisedResp())

		return
	}

	h.base.WriteJson(ctx, w, http.StatusOK, req.GetSingleItemResp(record))
}

func filterRules() []req.FilterRule {
	return []req.FilterRule{
		{
			Field: "deletedAt",
			Type:  req.FilterIsNull,
		},
		{
			Field: "name",
			Type:  req.FilterLike,
		},
	}
}

// getSortableFields will return a list of fields that a collection of records can be sorted by. This is necessary because
// not all fields should this be available to, and it will prevent SQL injection.
func getSortableFields() []string {
	return []string{
		"id",
		"name",
		"createdAt",
		"updatedAt",
	}
}

func (h *Handle) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	h.base.Log.Info(ctx, "Request made to Get myNewTable", nil)

	p := req.GetRequestParams(r, filterRules(), getSortableFields())

	records, err := h.srv.GetAll(ctx, p)
	if err != nil {
		h.base.WriteJson(ctx, w, http.StatusInternalServerError, "Unable to process request")

		return
	}

	h.base.WriteJson(ctx, w, http.StatusOK, req.GetListResp(records, p))
}
