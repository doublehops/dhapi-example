package model

import (
	"github.com/doublehops/dhapi-example/internal/model"
	req "github.com/doublehops/dhapi-example/internal/request"
	"github.com/doublehops/dhapi-example/internal/validator"
)

type MyNewTable struct {
	model.BaseModel
	Name       string `json:"name"`
	ID         int    `json:"id"`
	CurrencyID int    `json:"currencyId"`
}

func (m *MyNewTable) getRules() []validator.Rule {
	return []validator.Rule{
		{"name", m.Name, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}},
		{"id", m.ID, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}},
		{"currencyId", m.CurrencyID, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}},
	}
}

func (m *MyNewTable) Validate() req.ErrMsgs {
	return validator.RunValidation(m.getRules())
}
