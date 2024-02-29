package model

import (
	"github.com/doublehops/dhapi-example/internal/validator"

	req "github.com/doublehops/dhapi-example/internal/request"
)

type Author struct {
	BaseModel
	Name string `json:"name"`
}

func (a *Author) getRules() []validator.Rule {
	return []validator.Rule{
		{"name", a.Name, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}}, //nolint:govet
	}
}

func (a *Author) Validate() req.ErrMsgs {
	return validator.RunValidation(a.getRules())
}
