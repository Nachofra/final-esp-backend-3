package query_builder

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

type Expression struct {
	Key   string
	Exp   string
	Value interface{}
}

func (e Expression) ToString() string {

	switch e.Value.(type) {

	case int, int16, int32, int64:
		val := strconv.Itoa(e.Value.(int))
		clause := e.Key + e.Exp + e.getReplaceExp()
		return fmt.Sprintf(clause, val)

	default:
		if strings.TrimSpace(e.Value.(string)) == "" {
			return ""
		} else {
			e.Value = template.HTMLEscapeString(e.Value.(string))
			clause := e.Key + e.Exp + e.getReplaceExp()
			val := fmt.Sprintf(clause, e.Value)
			return val
		}
	}
}

func (e Expression) getReplaceExp() string {
	switch e.Value.(type) {
	case int, int64, int32, int16:
		return "%s"
	default:
		return "'%s'"
	}
}
