package request

type FilterType string

var FilterEquals FilterType
var FilterLike FilterType
var FilterIsNull FilterType
var FilterRange FilterType

type FilterRule struct {
	Field string
	Type  FilterType
	Value any
}

//type WhereClause struct {
//	Definition string
//	Value      any
//}
