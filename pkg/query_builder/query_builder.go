package query_builder

import (
	"strconv"
	"strings"
)

// DynamicQueryBuilder is the core of the package, it wraps a native string type.
type DynamicQueryBuilder string

// NewExpression creates a new SQL Expression.
func (dqb DynamicQueryBuilder) NewExpression(key string, assignment string, value interface{}) Expression {
	return Expression{Key: key, Exp: assignment, Value: value}
}

// And wraps multiple expressions with the "AND" operator of SQL.
func (dqb DynamicQueryBuilder) And(component ...interface{}) DynamicQueryBuilder {
	return dqb.getOperationExpression("AND", component...)
}

// Or wraps multiple expressions with the "OR" operator of SQL.
func (dqb DynamicQueryBuilder) Or(component ...interface{}) DynamicQueryBuilder {
	return dqb.getOperationExpression("OR", component...)
}

// getOperationExpression generates a DynamicQueryBuilder by combining multiple expressions with the specified SQL operator.
// Example of SQL operators: "AND", "OR".
func (dqb DynamicQueryBuilder) getOperationExpression(operation string, component ...interface{}) DynamicQueryBuilder {
	if len(component) == 0 {
		return ""
	}
	if len(component) == 1 {
		return componentToString(component[0])
	} else {
		clauses := make([]string, 0)
		for _, v := range component {
			value := componentToString(v)
			if value != "" {
				clauses = append(clauses, ""+string(value)+"")
			}
		}

		if len(clauses) > 0 {
			return DynamicQueryBuilder("( " + strings.Join(clauses, " "+operation+" ") + ")")
		}

		return ""
	}
}

// Limit sets "LIMIT" and "OFFSET" for the query.
func (dqb DynamicQueryBuilder) Limit(offset int, length int) DynamicQueryBuilder {
	query := string(dqb)
	query += " LIMIT " + strconv.Itoa(length) + " OFFSET " + strconv.Itoa(offset)
	return DynamicQueryBuilder(query)
}

// CopyQuery parses a string query into the DynamicQueryBuilder type and stores the result in the provided destination string.
// It returns the DynamicQueryBuilder for method chaining if required.
func (dqb DynamicQueryBuilder) CopyQuery(dest *string) DynamicQueryBuilder {
	*dest = dqb.ToString()
	return dqb
}

// BindSql generates and returns the final SQL query by combining the DynamicQueryBuilder with an existing SQL string.
// Example of existing SQL string: "SELECT * FROM table".
func (dqb DynamicQueryBuilder) BindSql(sql string) string {
	if dqb != "" && dqb != "( )" {
		index := strings.Index(dqb.ToString(), "LIMIT")
		if index == 1 {
			return sql + dqb.ToString()
		}

		return sql + " WHERE " + string(dqb)
	}
	return sql
}

// ToString converts the DynamicQueryBuilder to a string representation of the query.
func (dqb DynamicQueryBuilder) ToString() string {
	return string(dqb)
}

// componentToString converts various types of components into a DynamicQueryBuilder for query construction.
func componentToString(c interface{}) DynamicQueryBuilder {
	switch v := c.(type) {
	case Expression:
		return DynamicQueryBuilder(c.(Expression).ToString())
	case string, *string:
		return DynamicQueryBuilder(c.(string))
	case DynamicQueryBuilder:
		return v
	default:
		return ""
	}
}
