package model

import (
	"github.com/doublehops/dhapi-example/internal/model"
	req "github.com/doublehops/dhapi-example/internal/request"
	"github.com/doublehops/dhapi-example/internal/validator"
)

type MyNewTable struct {
	model.BaseModel
	CurrencyID int    `json:"currencyId"`
	Name       string `json:"name"`
}

func (m *MyNewTable) getRules() []validator.Rule {
	return []validator.Rule{
		// nolint:composites
		{"currencyId", m.CurrencyID, true, []validator.ValidationFuncs{validator.IsInt("")}},
		{"name", m.Name, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}},
	}
}

func (m *MyNewTable) Validate() req.ErrMsgs {
	return validator.RunValidation(m.getRules())
}
