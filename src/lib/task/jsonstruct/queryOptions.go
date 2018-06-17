package jsonstruct

import sq "github.com/Masterminds/squirrel"

//FilterOptions : query options
type FilterOptions struct {
	Column    string
	Value     string
	Operation string
}

//QueryOptions : query options
type QueryOptions struct {
	Command string
	Columns []string
	Filter  sq.Eq
	Table   string
}
