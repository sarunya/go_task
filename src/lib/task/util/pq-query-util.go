package util

import (
	"fmt"
	"strings"

	"../jsonstruct"

	sq "github.com/Masterminds/squirrel"
)

//ConstructQuery : constructs query
func ConstructQuery(queryOption jsonstruct.QueryOptions) string {
	var querystr string
	switch command := queryOption.Command; command {
	case "select":
		querystr = constructSelectQueryString(queryOption)

	}
	return querystr
}

//ConstructQueryBuilder : constructs query
func ConstructQueryBuilder(queryOption jsonstruct.QueryOptions) sq.SelectBuilder {
	var querySlctr sq.SelectBuilder
	switch command := queryOption.Command; command {
	case "select":
		querySlctr = constructSelectQuery(queryOption)

	}
	return querySlctr
}

func constructSelectQueryString(queryOption jsonstruct.QueryOptions) string {
	columns := strings.Join(queryOption.Columns, ",")
	query := sq.Select(columns).From(queryOption.Table)
	active := query.Where(queryOption.Filter)
	sql, args, err := active.ToSql()
	fmt.Print(args, err, sql)
	return sql
}

func constructSelectQuery(queryOption jsonstruct.QueryOptions) sq.SelectBuilder {
	columns := strings.Join(queryOption.Columns, ",")
	query := sq.Select(columns).From(queryOption.Table)
	active := query.Where(queryOption.Filter)
	return active
}
