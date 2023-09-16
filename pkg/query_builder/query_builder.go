package query_builder

import (
	"strconv"
	"strings"
)

type DynamicQueryBuilder string

func (dqb DynamicQueryBuilder) NewExpression(key string, assignment string, value interface{}) Expression {
	return Expression{Key: key, Exp: assignment, Value: value}
}

func (dqb DynamicQueryBuilder) And(component ...interface{}) DynamicQueryBuilder {
	return dqb.getOperationExpression("AND", component...)
}

func (dqb DynamicQueryBuilder) Or(component ...interface{}) DynamicQueryBuilder {
	return dqb.getOperationExpression("OR", component...)
}

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

func (dqb DynamicQueryBuilder) Limit(offset int, length int) DynamicQueryBuilder {
	query := string(dqb)
	query += " LIMIT " + strconv.Itoa(length) + " OFFSET " + strconv.Itoa(offset)
	return DynamicQueryBuilder(query)
}

func (dqb DynamicQueryBuilder) CopyQuery(dest *string) DynamicQueryBuilder {
	*dest = dqb.ToString()
	return dqb
}

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

func (dqb DynamicQueryBuilder) ToString() string {
	return string(dqb)
}

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
