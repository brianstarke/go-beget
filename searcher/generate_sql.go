package searcher

import (
	"fmt"
	"log"

	"github.com/lann/squirrel"
	"github.com/mgutz/ansi"
)

var (
	logPrefix = "[" +
		ansi.Color("go-beget", "154") +
		"/" +
		ansi.Color("searcher", "159") + "] "
)

// GenerateSelectSQL will generate a SQL search query from a SearchRequest.
//
// This returns the SQL string with placeholders and the placeholder values
// with the idea that you're running the SQL something like:
//
//   var sr search.ThingSearchRequest
//
//   sr.AddFilter(search.ThingColor, "RED", searcher.EQ, searcher.AND)
//   sr.Limit = 5
//
//   sqlStr, values, _ := searcher.GenerateSelectSQL(sr)
//
//   t := []Thing{}
//   err := db.Select(&t, sqlStr, values)
//
// (example uses sqlx http://jmoiron.github.io/sqlx/#getAndSelect)
// NOTE: I've only tried this with PostgreSQL
func GenerateSelectSQL(sr SearchRequest) (sqlStr string, values interface{}, err error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	var sql squirrel.SelectBuilder

	if len(sr.GetFields()) == 0 {
		sql = psql.Select("*")
	} else {
		sql = psql.Select(sr.GetFields()...)
	}

	sql = sql.From(sr.GetTableName())

	getPredicates(&sql, sr.GetFilters())

	if sr.GetLimit() > 0 {
		sql = sql.Limit(uint64(sr.GetLimit()))
	}

	if sr.GetOffset() > 0 {
		sql = sql.Offset(uint64(sr.GetOffset()))
	}

	if sr.GetOrderBy().Field.DbFieldName() != "" {
		if sr.GetOrderBy().Descending {
			sql = sql.OrderBy(sr.GetOrderBy().Field.DbFieldName() + " DESC")
		} else {
			sql = sql.OrderBy(sr.GetOrderBy().Field.DbFieldName())
		}
	}

	sqlStr, values, err = sql.ToSql()

	log.SetFlags(0)
	log.SetPrefix(logPrefix)
	log.Printf("SELECT SQL generated %s", ansi.Color(sqlStr, "155+b"))
	log.SetPrefix("")

	return sqlStr, values, err
}

func getPredicates(sq *squirrel.SelectBuilder, filters []Filter) {
	ors := getOrs(filters)
	ands := getAnds(filters)

	if len(ors) > 0 {
		*sq = sq.Where(ors)
	}

	if len(ands) > 0 {
		*sq = sq.Where(ands)
	}
}

func getOrs(filters []Filter) squirrel.Or {
	ors := squirrel.Or{}

	for _, f := range filters {
		if f.Condition == OR {
			if f.Operator == ISNULL || f.Operator == ISNOTNULL {
				ors = append(ors, squirrel.Expr(buildExpr(f)))
			} else {
				ors = append(ors, squirrel.Expr(buildExpr(f), f.Value))
			}
		}
	}

	return ors
}

/*
AND is default if no condition is provided
*/
func getAnds(filters []Filter) squirrel.And {
	ands := squirrel.And{}

	for _, f := range filters {
		if f.Condition == AND || f.Condition == "" {
			if f.Operator == ISNULL || f.Operator == ISNOTNULL {
				ands = append(ands, squirrel.Expr(buildExpr(f)))
			} else {
				ands = append(ands, squirrel.Expr(buildExpr(f), f.Value))
			}
		}
	}

	return ands
}

func buildExpr(f Filter) string {
	var operator string

	// default to EQ if not provided
	if f.Operator == "" || f.Operator == EQ {
		operator = "="
	} else {
		switch f.Operator {
		case EQ:
			operator = "="
		case IEQ:
			operator = "ILIKE"
		case GT:
			operator = ">"
		case GTE:
			operator = ">="
		case LT:
			operator = "<"
		case LTE:
			operator = "<="
		case NE:
			operator = "!="
		case NIE:
			operator = "NOT ILIKE"
		case LIKE:
			operator = "ILIKE"
		case ISNULL:
			operator = "IS NULL"
		case ISNOTNULL:
			operator = "IS NOT NULL"
		}
	}

	if f.Operator == ISNULL || f.Operator == ISNOTNULL {
		return fmt.Sprintf("%s %s", string(f.Field.DbFieldName()), operator)
	}

	if f.Operator == LIKE {
		return fmt.Sprintf("cast (%s as varchar(64)) %s ?", string(f.Field.DbFieldName()), operator)
	}
	return fmt.Sprintf("%s %s ?", string(f.Field.DbFieldName()), operator)
}
