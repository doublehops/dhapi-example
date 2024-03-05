package model

import (
	"github.com/doublehops/dh-go-framework/internal/model"
	req "github.com/doublehops/dh-go-framework/internal/request"
	"github.com/doublehops/dh-go-framework/internal/validator"
)

type MyNewTable struct {
	model.BaseModel
	CurrencyID int    `json:"currencyId"`
	Name       string `json:"name"`
}

func (m *MyNewTable) getRules() []validator.Rule {
	return []validator.Rule{
		{"currencyId", m.CurrencyID, true, []validator.ValidationFuncs{validator.IsInt("")}},   //nolint:govet
		{"name", m.Name, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}}, //nolint:govet

	}
}

func (m *MyNewTable) Validate() req.ErrMsgs {
	return validator.RunValidation(m.getRules())
}
