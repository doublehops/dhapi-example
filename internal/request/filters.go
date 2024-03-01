package request

type FilterType string

var (
	FilterEquals FilterType = "FilterEquals"
	FilterLike   FilterType = "FilterLike"
	FilterIsNull FilterType = "FilterIsNull"
	FilterRange  FilterType = "FilterRange"
)

type FilterRule struct {
	Field string
	Type  FilterType
	Value any
}

type Params []any
