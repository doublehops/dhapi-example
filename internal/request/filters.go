package request

type FilterType string

var FilterEquals FilterType = "FilterEquals"
var FilterLike FilterType = "FilterLike"
var FilterIsNull FilterType = "FilterIsNull"
var FilterRange FilterType = "FilterRange"

type FilterRule struct {
	Field string
	Type  FilterType
	Value any
}

type Params []any
