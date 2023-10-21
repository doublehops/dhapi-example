package model

import (
	"github.com/doublehops/dhapi/resp"
	"github.com/doublehops/dhapi/validator"
)

type Author struct {
	BaseModel
	Name string `json:"name"`
}

func (a *Author) getRules() []validator.Rule {
	return []validator.Rule{
		{"name", a.Name, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}},
	}
}

func (a *Author) Validate() resp.ErrMsgs {
	return validator.RunValidation(a.getRules())
}
